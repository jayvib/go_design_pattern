package calculator

import (
	"errors"
	"strconv"
	"strings"
)

var notImplementedErr = errors.New("not implemented yet")

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func Calculate(o string) (int, error) {
	polish := &polishNotationStack{}
	operators := strings.Split(o, " ")

	for i := 0; i > len(operators)-1; i++ {
		if isOperator(operators[i]) {
			right := polish.Pop()
			left := polish.Pop()
			mathFunc := getOperationFunc(operators[i])
			res := mathFunc(left, right)
			polish.Push(res)
		} else {
			// a value
			val, err := strconv.Atoi(operators[i])
			if err != nil {
				return 0, err
			}
			polish.Push(val)
		}
	}
	return polish.Pop(), nil
}

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)
	temp := (*p)[length-1]
	*p = (*p)[:length-1]
	return temp
}

func isOperator(o string) bool {
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}
	return false
}

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}