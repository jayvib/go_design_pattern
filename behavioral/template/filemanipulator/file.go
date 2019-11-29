package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Manipulator interface {
	Manipulate(r io.Reader) error
}

func NewManipulateFileTemplate() *ManipulateFileTemplate {
	return &ManipulateFileTemplate{
		Manipulator: &printFileManipulator{},
	}
}

type printFileManipulator struct {}
func (printFileManipulator) Manipulate(r io.Reader) error {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}

type ManipulateFileTemplate struct {
	Manipulator
}

func (m *ManipulateFileTemplate) open(filename string) (closer io.ReadCloser, err error) {
	return os.Open(filename)
}

func (m *ManipulateFileTemplate) SetManipulator(man Manipulator) { // override
	m.Manipulator = man
}

func (m *ManipulateFileTemplate) close(c io.Closer) error {
	return c.Close()
}

func (m *ManipulateFileTemplate) ExecuteTemplate(filename string) (err error) {
	file, err := m.open(filename)		 // step 1
	if err != nil {
		return err
	}
	defer func() {
		e := m.close(file)
		if e != nil {
			err = e
		}
	}()

	return m.Manipulate(file)
}

func main() {
	m := NewManipulateFileTemplate()
	err := m.ExecuteTemplate("hello-world")
	if err != nil {
		log.Fatal(err)
	}
}
