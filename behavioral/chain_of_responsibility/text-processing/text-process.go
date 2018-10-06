package text_processing

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

func (u *ToUpperCase) TransformNext(s string) {

}

type ToTitle struct {}

func (t *ToTitle) TransformNext(s string) {

}

type ToReplace struct {
	charToReplace string
	charToReplaceWith string
}

func (r *ToReplace) TransformNext(s string) {

}

