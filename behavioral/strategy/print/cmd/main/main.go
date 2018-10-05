package main

import (
	"flag"
	print2 "go_design_pattern/behavioral/strategy/print"
	"log"
	"os"
)

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()
	activeStrategy, err := print2.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case print2.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case print2.IMAGE_STRATEGY:
		w, err := os.Create("./image/image.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
