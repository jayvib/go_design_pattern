package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jayvib/app"
)

func main() {
	http.Handle("/", app.NewLogMiddleware(os.Stdout,
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(res, "Hello World")
		}),
	))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
