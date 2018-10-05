package print

import (
	"fmt"
	"go_design_pattern/behavioral/strategy/print/shapes"
	"io"
	"os"
)

const (
	TEXT_STRATEGY = "text"
	IMAGE_STRATEGY = "image"
)

type Strategy interface {
	Print() error
}

type StrategyV2 interface {
	Strategy
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutput struct {
	Writer io.Writer
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

func (d *PrintOutput) Write(b []byte) (int, error) {
	return d.Writer.Write(b)
}

func (d *PrintOutput) LogWrite(b []byte) (int, error) {
	return d.LogWriter.Write(b)
}

func NewPrinter(s string) (StrategyV2, error) {
	switch s {
	case TEXT_STRATEGY:
		return &shapes.ConsoleSquare{
			Output: &PrintOutput{
				LogWriter: os.Stdout, // default out
			},
		}, nil
	case IMAGE_STRATEGY:
		return &shapes.ImageSquare{
			Output: &PrintOutput{
				LogWriter: os.Stdout, // default out
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
