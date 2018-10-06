package shapes

import "io"

type Output interface {
	SetLog(w io.Writer)
	SetWriter(w io.Writer)
	LogWrite(b []byte) (n int, err error)
	io.Writer
}

// PrintOutput satisfies Output interface.
//
// io.Writer is embedded in order to satisfy the Write method in Output interface.
type PrintOutput struct {
	io.Writer // embedd the object that implements a router so that this object will inherit its method
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

func (d *PrintOutput) LogWrite(b []byte) (int, error) {
	return d.LogWriter.Write(b)
}
