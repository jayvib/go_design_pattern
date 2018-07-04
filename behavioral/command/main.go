package main

import (
	"fmt"
	"go_design_pattern/behavioral/command/task"
)

type task1 struct {
	name string
}

func (t *task1) Do() {
	fmt.Printf("Doing task1 by %s\n", t.name)
}

type taskFunc func()

func (tf taskFunc) Do() {
	tf()
}

type taskFunc2 struct {
	name  string
	items []string
	fn    func(string, ...string)
}

func (tf2 *taskFunc2) Do() {
	tf2.fn(tf2.name, tf2.items...)
}

func taskFuncHelper(name string, fn func(string, ...string), items ...string) task.Command {
	var cmd task.Command
	cmd = &taskFunc2{
		name:  name,
		items: items,
		fn:    fn,
	}
	return cmd
}

func main() {

	t1 := &task1{"Jayson"}
	tfunc := taskFunc(func() {
		msg := "I'm in hurry, don't disturbe me while I'm doing this task okay?"
		fmt.Println(msg)
	})

	ts := task.New(t1, tfunc)

	tf2 := taskFuncHelper("Foo", func(n string, items ...string) {
		msg := fmt.Sprintf("%s is collecting some:", n)
		for _, i := range items {
			msg += " " + i
		}
		fmt.Println(msg)
	}, "apple", "orange", "grapes", "banana")

	ts.AddTask(tf2)

	ts.Execute()
}
