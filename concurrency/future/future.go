package future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc FailFunc
}

func (m *MaybeString) Success(fn SuccessFunc)  *MaybeString {
	m.successFunc = fn
	return m
}

func (m *MaybeString) Fail(fn FailFunc) *MaybeString {
	m.failFunc = fn
	return m
}

func (m *MaybeString) Execute(fn ExecuteFunc) {
	go func(m *MaybeString) {
		str, err := fn()
		if err != nil {
			m.failFunc(err)
		} else {
			m.successFunc(str)
		}
	}(m)
}