package publisher_subscriber

import (
	"errors"
	"io"
)

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	return &writerSubscriber{
		id:     id,
		Writer: out,
	}
}

type writerSubscriber struct {
	id     int
	Writer io.Writer
}

func (s *writerSubscriber) Notify(msg interface{}) error {
	return errors.New("not implemented yet")
}

func (s *writerSubscriber) Close() {}
