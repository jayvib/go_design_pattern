package display

// Define an interface to be used for bridge design pattern

type DisplayImpl interface {
	rawOpen() string
	rawPrint() string
	rawClose() string
}

type StringDisplayImpl struct {
	str string
}

func (self *StringDisplayImpl) rawOpen() string {
	return self.printLine()
}

func (self *StringDisplayImpl) rawPrint() string {
	return "|" + self.str + "|\n"
}

func (self *StringDisplayImpl) rawClose() string {
	return self.printLine()
}

func (self *StringDisplayImpl) printLine() string {
	str := "+"

	for _, _ = range self.str {
		str += "-"
	}

	str += "+"
	return str
}

