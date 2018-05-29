package main

import (
	"io"
	"os"
	"strconv"
)

type Counter struct {
	Writer io.Writer
}

func (c *Counter) Count(n uint64) uint64 {
	if n == 0 {
		c.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	c.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return c.Count(n - 1)
}

func main() {
	file, _ := os.Create("out.txt")
	defer file.Close()

	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()

	counter := Counter{Writer: pipeWriter}
	counter.Count(5)

	io.TeeReader(pipeReader, file)

	//go func() {
	//	io.Copy(os.Stdout, tee)
	//}()

}
