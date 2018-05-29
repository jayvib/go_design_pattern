package display

type DefaultDisplay struct {
	impl DisplayImpl
}

func (dd *DefaultDisplay) open() string {
	return dd.impl.rawOpen()
}

func (dd *DefaultDisplay) print() string {
	return dd.impl.rawPrint()
}

func (dd *DefaultDisplay) close() string {
	return dd.impl.rawClose()
}

func (dd *DefaultDisplay) Display() string {
	return dd.open() + dd.print() + dd.close()
}
