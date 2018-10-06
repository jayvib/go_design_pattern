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


func NewPrinter(s string) (StrategyV2, error) {
	switch s {
	case TEXT_STRATEGY:
		return &shapes.ConsoleSquare{
			Output: &shapes.PrintOutput{
				LogWriter: os.Stdout, // default out
			},
		}, nil
	case IMAGE_STRATEGY:
		return &shapes.ImageSquare{
			Output: &shapes.PrintOutput{
				LogWriter: os.Stdout, // default out
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
