package printer

import (
	"testing"
	"fmt"
)

func TestTemplatePrinter(t *testing.T) {
	printer := &UserPrinter{content: "Hello Jayson"}

	absPt := NewAbstractPrinter(printer)

	res := absPt.Display()
	fmt.Println(res)
}
