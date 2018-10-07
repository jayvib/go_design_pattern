package text_processing

import (
	"fmt"
	"strings"
)

// Criteria:
// 1. An object that will implement the chain text processor interface.

type ChainTextProcessor interface {
	TransformNext(s string)
}

func NewTextProcessor(textProcessor ...ChainTextProcessor) *TextProcessor {
	tp := &TextProcessor{
		chainTextProcessors: make([]ChainTextProcessor, 0),
	}

	for _, c := range textProcessor {
		tp.chainTextProcessors = append(tp.chainTextProcessors, c)
	}

	return tp
}

type TextProcessor struct {
	chainTextProcessors []ChainTextProcessor
}

func (p *TextProcessor) Transform(s string) {
	for _, c := range p.chainTextProcessors {
		c.TransformNext(s)
	}
}

type ToUpperCase struct {}

func (ToUpperCase) TransformNext(s string) {
	fmt.Println(strings.ToUpper(s))
}

type ToLower struct {}

func (ToLower) TransformNext(s string) {
	fmt.Println(strings.ToLower(s))
}

type ToReplace struct {
	charToReplace string
	charToReplaceWith string
}

func (r *ToReplace) TransformNext(s string) {
	fmt.Println(strings.Replace(s, r.charToReplace, r.charToReplaceWith, -1))
}

