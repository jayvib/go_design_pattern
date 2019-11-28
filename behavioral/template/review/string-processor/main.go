package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type processorFunc func(string) string

type template struct {
	msg string
}

func (t template) addprefix(str string) string {
	return "-" + str
}

func (t template) addsuffix(str string) string {
	return str + "-"
}

func (t template) Execute(fn processorFunc) string {
	return t.addsuffix(fn(t.addprefix(t.msg)))

}

func main() {
	t := template{msg: "Hello World"}
	res := t.Execute(func(str string) string {
		return strings.ReplaceAll(str, " ", "|")
	})
	logrus.Println(res)
}
