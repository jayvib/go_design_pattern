package main

import (
	"flag"
	"log"
	"os"
)

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()
	activeStrategy, err := print.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case print.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case print.IMAGE_STRATEGY:
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
