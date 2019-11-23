package birthday_party

func NewBirthdayEvent(day string) *BirthdayEvent {
	return &BirthdayEvent{
		Day: day,
	}
}

type BirthdayEvent struct {
	Day      string
	Attendee []string
}

func (b *BirthdayEvent) AddAttendee(a ...string) {
	for _, n := range a {
		b.Attendee = append(b.Attendee, n)
	}
}

func (b *BirthdayEvent) GetName() []string {
	return b.Attendee
}

func (b *BirthdayEvent) Accept(v Visitor) {
	v.Visit(b)
}
