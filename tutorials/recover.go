package main

import "fmt"

//go makes it possible to recover from a panic by using the built in recover function
//a recover can stop a panic from aborting the program and let it continue with exectuoin instead
//an example when to use? When a server doesn't want to crash if one of hte client connections exhibits a critical error instead just close that connection and continue servign other clients
//go net/http does it by default for http servers

func mayPanic() {
	panic("a problem")
}

func main()  {

	//recover must be called within a defer function.
	//when enclosing function panics defer activates and a recover call within it catches the panic
	//return value of recover is the error raised in the panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}