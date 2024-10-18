package main

import (
	"fmt"
	"time"
)

//can use channels to synchronize execution across goroutines
//when waiting for multiple goroutines to finish may prefer to use a waitgroup

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	//start the worker go routine and give it a channel to notify on it
	go worker(done)

	//block until we receive a notifiation from the worker on the channel
	//without this line the program would exit before the worker even started
	//this is a blocking receive to allow a goroutine to finish
	<-done
}