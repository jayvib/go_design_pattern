package database

import (
	"io"
	"fmt"
	"errors"
)

type fileStorage struct {
	file io.ReadWriteCloser
}

func (f *fileStorage) Put(key, value string) error {
	_, err :=  f.file.Write([]byte(fmt.Sprintf("%s|%s\n", key, value)))
	return err
}

func (f *fileStorage) Post(key, value string) error {
	return errors.New("not implemented yet")
}

func (f *fileStorage) Get(key string) (value string, err error) {
	return "", errors.New("not implemented yet")
}

func (f *fileStorage) Delete(key string) error {
	return errors.New("not implemented yet")
}
