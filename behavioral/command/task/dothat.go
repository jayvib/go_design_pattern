package task

type Command interface {
	Do()
}

type TaskStorage struct {
	tasks []Command
}

func (ts *TaskStorage) AddTask(cmd ...Command) {
	ts.tasks = append(ts.tasks, cmd...)
}

func (ts *TaskStorage) Execute() {
	for _, cmd := range ts.tasks {
		cmd.Do()
	}
}

func New(cmd ...Command) *TaskStorage {
	ts := new(TaskStorage)

	for _, v := range cmd {
		ts.tasks = append(ts.tasks, v)
	}

	return ts
}
