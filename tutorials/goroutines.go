package main

import (
	"fmt"
	"time"
)

//goroutine being a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	//invoking f in a goroutine
	go f("goroutine")

	//starting a goroutine for an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//more proper way to wait for th goroutines is with a waitgroup
	time.Sleep(time.Second)
	fmt.Println("done")
}