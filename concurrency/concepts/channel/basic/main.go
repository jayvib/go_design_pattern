package main

import "time"

func main() {
	channel := make(chan string)

	go func() {
		println("goroutine will sleep")
		time.Sleep(3 * time.Second)
		channel <- "Hello Jayson!"
	}()

	x := <-channel // this will be block until the goroutine will send a data to the channel.
	println(x)
}
