package main

import "fmt"

func main() {

	//created a new channel with make(chan val-type)
	//channels are typed by the values they convey
	messages := make(chan string)

	// channel <- syntax sends a value into the channel. here we send ping to the messages channel from a new goroutine
	go func()  { messages <- "ping" } ()

	// <- channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)
}