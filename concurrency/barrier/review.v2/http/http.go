package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buff bytes.Buffer
		io.Copy(&buff, reader)
		out <- buff.String()
	}()

	barriers(endpoints...)

	writer.Close()
	tmp := <-out
	return tmp
}

type Response struct {
	Body string
	Err  error
}

func barriers(endpoints ...string) {
	requestNum := len(endpoints)

	inChan := make(chan Response, requestNum)
	defer close(inChan)
	responses := make([]Response, requestNum)

	for _, endpoint := range endpoints {
		go makeRequest(inChan, endpoint)
	}

	for i := 0; i < requestNum; i++ {
		resp := <-inChan
		responses = append(responses, resp)
	}

	for _, r := range responses {
		if r.Err == nil {
			fmt.Println(r.Body)
		}
	}
}

func makeRequest(in chan<- Response, endpoint string) {
	var res Response

	resp, err := http.Get(endpoint)
	if err != nil {
		res.Err = err
		in <- res
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		in <- res
		return
	}

	res.Body = string(content)
	in <- res
}
