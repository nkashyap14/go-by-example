package main

//use panics to fail fast on errors that shouldn't happen during normal operation

import "os"

func main() {
	//panic("a problem")

	//typically use pianic if a function returns an error value that we don't know how to or want to handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}