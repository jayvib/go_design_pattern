package main

import "fmt"

// https://refactoring.guru/design-patterns/command

type Saver interface {
	Save()
}

type Undoer interface {
	Undo()
}

type SaveUndoer interface {
	Saver
	Undoer
}

type Command interface {
	Execute()
}

func NewSaveCommand(saver Saver) Command {
	return &saveCommand{saver:saver}
}


// Command Pattern suggests that GUI objects should'nt send these requests directly.
// Instead, you should extract all of the requests details, such as the object being
// called, the name of the method and the list of arguments into a separate command class
// with a single method that triggers this request.
type saveCommand struct {
	saver Saver // strategy pattern
}

func (s *saveCommand) Execute() {
	s.saver.Save()
}

func NewUndoCommand(undoer Undoer) Command {
	return &undoCommand{undoer:undoer}
}

type undoCommand struct {
	undoer Undoer
}

func (u *undoCommand) Execute() {
	u.undoer.Undo()
}

func NewSaveUndoer() SaveUndoer {
	return saveUndoerImpl{}
}

type saveUndoerImpl struct {}

func (saveUndoerImpl) Save() {
	fmt.Println("Saving")
}

func (saveUndoerImpl) Undo() {
	fmt.Println("Undo")
}

func NewButton(name string, command Command) *Button {
	return &Button{Name:name, command:command}
}

type Button struct { // invoker
	Name string
	command Command
}

func (b *Button) Push() {
	fmt.Printf("Button %s pushed\n", b.Name)
	b.command.Execute()
}

type Shortcut struct {
	Name string
	command Command
}

func (s *Shortcut) Do() {
	s.command.Execute()
}

func main() {
	saverUndoer := NewSaveUndoer()
	saver := NewSaveCommand(saverUndoer)
	undoer := NewUndoCommand(saverUndoer)

	saveButton := NewButton("save", saver)
	undoerButton := NewButton("undo", undoer)

	saveButton.Push()
	undoerButton.Push()


}
