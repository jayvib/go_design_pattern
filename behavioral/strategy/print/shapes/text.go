package shapes

import (
	"bytes"
	"io"
)


type ConsoleSquare struct {
	Output
}

func (c *ConsoleSquare) Print() error {
	r := bytes.NewReader([]byte("Square"))
	_, err := io.Copy(c.Output, r)
	if err != nil {
		return err
	}
	return nil
}
