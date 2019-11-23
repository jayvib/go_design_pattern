package main

import (
	"github.com/jayvib/go_design_pattern/behavioral/command/command_chain_of_responsibility/time-command"
	"time"
)

func main() {
	second := new(time_command.Logger)
	first := time_command.Logger{
		NextChain: second,
	}
	command := &time_command.TimePassed{
		Start: time.Now(),
	}

	first.Next(command)
}
