package main

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFn SuccessFunc
	failFn    FailFunc
}

func (m *MaybeString) Success(f SuccessFunc) *MaybeString {
	m.successFn = f
	return m
}

func (m *MaybeString) Fail(f FailFunc) *MaybeString {
	m.failFn = f
	return m
}

func (m *MaybeString) Execute(fn ExecuteStringFunc) {
	go func(m *MaybeString) {
		s, err := fn()
		if err != nil {
			m.failFn(err)
			return
		}

		m.successFn(s)
	}(m)
}
