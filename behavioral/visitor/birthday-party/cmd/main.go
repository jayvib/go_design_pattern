package main

import "github.com/go_design_pattern/behavioral/visitor/birthday-party"

func main() {
	prty := birthday_party.NewBirthdayEvent("monday")
	prty.AddAttendee("jayson", "jimboy", "jumily", "geraldine")

	prty2 := birthday_party.NewBirthdayEvent("tuesday")
	prty.AddAttendee("jonora", "jesiel","jessica", "joseph", "jericho")

	visitor := &birthday_party.EventCoordinator{}

	prty.Accept(visitor)
	prty2.Accept(visitor)

	visitor.Print()
}


