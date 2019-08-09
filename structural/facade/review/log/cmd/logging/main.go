package main

import (
	"github.com/jayvib/go_design_pattern/structural/facade/review/log"
)

func main() {
	log.SetLogger(log.Logrus)
	log.Println("Hello Logging World!")
}
