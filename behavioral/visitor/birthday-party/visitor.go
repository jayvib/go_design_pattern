package birthday_party

import (
	"log"
	"os"
	"text/template"
)

type EventCoordinator struct {
	NameListAttendee []string
}

func (e *EventCoordinator) Visit(nr NameRetriever) {
	e.NameListAttendee = append(e.NameListAttendee, nr.GetName()...)
}

func (e *EventCoordinator) Print() {
	tmpl := `-----------------
Attendees:
{{range .}}	-{{.}}
{{end}}-----------------
`

	temp := template.New("attendee")
	t, err := temp.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, e.NameListAttendee); err != nil {
		log.Fatal(err)
	}

}
