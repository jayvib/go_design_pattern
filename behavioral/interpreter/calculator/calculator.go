package calculator

import "errors"

var notImplementedErr = errors.New("not implemented yet")

func Calculate(o string) (int, error) {
	return 0, notImplementedErr
}

