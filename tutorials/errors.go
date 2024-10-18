package main

import (
	"errors"
	"fmt"
)

//by convention errors are the last return value and have type error a built in interface
func f(arg int) (int, error) {
	if arg == 42  {
		return -1, errors.New("Can't work with 42")
	}

	//nil is no error
	return arg + 3, nil
}


//sentinel errors: predeclared values that are used to signify a specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

//are wrapping errors with fmt.Errorf and the %w verb to add context. Wrapper errors create a logical chain that can be queried iwth funcitons like errors.Is and errors.As
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("Making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed", e)
		} else {
			fmt.Println("f worked", r)
		}
	}

	//error.sis checks that a given error matches a specific error value
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("Unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}