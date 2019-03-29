package main

type StringFutureExecuter interface {
	RegisterFailFunc(failFunc FailFunc) StringFutureExecuter
	RegisterSuccessFunc(successFunc SuccessFunc) StringFutureExecuter
	Execute(executeFunc ExecuteFunc)
}

type FailFunc func(error)
type SuccessFunc func(string)
type ExecuteFunc func() (string, error)

type stringFuture struct {
	FailFunc
	SuccessFunc
	ExecuteFunc
}

func (s *stringFuture) RegisterFailFunc(fn FailFunc) StringFutureExecuter {
	return s
}

func (s *stringFuture) RegisterSuccessFunc(fn SuccessFunc) StringFutureExecuter {
	return s
}

func (s *stringFuture) Execute(executeFunc ExecuteFunc) {
	res, err := executeFunc()
	if err != nil {
		s.FailFunc(err)
	} else {
		s.SuccessFunc(res)
	}
}
