package where_to_write

import (
	"errors"
	"io"
)

type method int32

const (
	text method = iota
	csv
)

type WriteMethod interface {
	io.WriteCloser
}

func GetWriteMethod(m method) (WriteMethod, error) {
	return nil, errors.New("Not implemented yet")
}

type Text struct {
}
