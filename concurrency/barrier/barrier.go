package barrier

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type barrierResp struct {
	Err  error
	Resp string
}

var timeoutMillisecond = 5000

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer // substitute the writer that was got from the pipe

	out := make(chan string)
	go func() {
		var buf bytes.Buffer  // store the content to the buffer
		io.Copy(&buf, reader) // this will block until the reader has something to read
		out <- buf.String()   // send the content to the out channel
	}()

	barrier(endpoints...)
	writer.Close()
	temp := <-out
	return temp
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber) // a channel for barrierResp with expected number of request.
	defer close(in)

	responses := make([]barrierResp, requestNumber) // store the responses

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint) // deploy goroutines to download
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMillisecond) * time.Millisecond),
	}
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}
