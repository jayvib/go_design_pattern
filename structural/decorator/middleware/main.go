package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

type myServer struct {}

func (myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Middleware!")
}

type LoggerServer struct {
	Handler http.Handler
	LogWriter io.Writer
}

func (l *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(l.LogWriter, "Request URI: %s\n", r.RequestURI)
	r.BasicAuth()
}

func main() {
	http.Handle("/", &LoggerServer{
		Handler: myServer{},
		LogWriter: os.Stdout,
	})
	http.ListenAndServe(":8080", nil)
}
