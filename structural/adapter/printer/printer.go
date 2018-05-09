package printer

import "fmt"

type ModernPrintable interface {
	printMessage() string
}

type OldPrintable interface {
	print() string
}

type LegacyPrinter struct {
	str string
}

func (lp *LegacyPrinter) SetMessage(str string) {
	lp.str = str
}

func (lp *LegacyPrinter) GetMessage() string {
	return lp.str
}

func (lp *LegacyPrinter) print() string {
	return "old+++" + lp.str + "+++old"
}

type ModernPrinter struct {
	str string
	oldPrintable OldPrintable
}

func (mp *ModernPrinter) SetOldPrinter(printer OldPrintable) {
	mp.oldPrintable = printer
}

func (mp *ModernPrinter) printMessage() string {
	msg := mp.oldPrintable.print()
	return mp.decorate(msg)
}

func (mp *ModernPrinter) decorate(str string) string {
	msg := fmt.Sprintf("--|modern|%s|modern|--", str)
	return msg
}


type SuperModernPrinter struct {
	modernPrinter ModernPrintable
}

func (self *SuperModernPrinter) SetModernPrinter(printer ModernPrintable) {
	self.modernPrinter = printer
}

func (self *SuperModernPrinter) Display() {
	fmt.Println(self.modernPrinter.printMessage())
}

