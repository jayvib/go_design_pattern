package main

import (
	"fmt"
	"sync"
)

func NewStringRequest(s string, id int, wg *sync.WaitGroup) *Request {
	return &Request{
		Data: s,
		Handler: func(i interface{}) {
			defer wg.Done()
			s := i.(string)
			fmt.Println(s)
		},
	}
}

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})
