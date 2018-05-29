package printer

import (
	"fmt"
	"testing"
)

func TestTemplatePrinter(t *testing.T) {
	printer := &UserPrinter{content: "Hello Jayson"}

	absPt := NewAbstractPrinter(printer)

	res := absPt.Display()
	fmt.Println(res)
}
