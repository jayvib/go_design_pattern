package worker

type Request struct {
	Data    interface{}
}

type Result struct {
	Request
	Res interface{}
	Err error
}
