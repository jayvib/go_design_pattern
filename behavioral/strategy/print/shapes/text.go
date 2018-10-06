package shapes

import (
	"bytes"
	"io"
)

type ConsoleSquare struct {
	Output // the PrintOutput will satisfy this interface
}

func (c *ConsoleSquare) Print() error {
	r := bytes.NewReader([]byte("Square"))
	_, err := io.Copy(c.Output, r) // since Output is a writer because the io.Writer is embedded to the interface.
	if err != nil {
		return err
	}
	return nil
}
