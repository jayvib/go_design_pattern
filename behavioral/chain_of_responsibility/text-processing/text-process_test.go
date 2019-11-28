package text_processing

import (
	"io"
	"testing"
)

func TestTextProcessing_ToUpperCase(t *testing.T) {
	var toUpperCase ToUpperCase

	toUpperCase.TransformNext("hello")
}

func TestTextProcessing_ToLowerCase(t *testing.T) {
	var toLowerCase ToLower
	toLowerCase.TransformNext("HELLO")
}

func TestTextProcessing_TextReplace(t *testing.T) {
	toReplace := &ToReplace{
		charToReplace:     "H",
		charToReplaceWith: "h",
	}
	toReplace.TransformNext("Hello")
}

func TestChainTextProcessing(t *testing.T) {
	textProcessor := NewTextProcessor(
		ToLower{},
		ToUpperCase{},
		&ToReplace{
			charToReplaceWith: "J",
			charToReplace:     "h",
		})

	textProcessor.Transform("hEllO")
	io.MultiWriter()
}
