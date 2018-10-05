package shapes

import "io"

type Output interface {
	SetLog(w io.Writer)
	SetWriter(w io.Writer)
	LogWrite(b []byte) (n int, err error)
	io.Writer
}
