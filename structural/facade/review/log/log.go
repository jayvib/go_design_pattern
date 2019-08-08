package log

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Println(args ...interface{})
}

var logger Logger

var (
	// This is a form of Facade
	Logrus = logrus.New()
	Std    = log.New(os.Stdout, "Standard:", log.Ldate|log.Ltime)
	
	// I can add more Logger implementations here
)

func SetLogger(l Logger) {
	logger = l
}

func Println(args ...interface{}) {
	logger.Println(args)
}
