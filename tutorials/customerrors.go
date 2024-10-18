package main

import (
	"errors"
	"fmt"
)

//can use custom type as errors by implmeneting Error() method on them
//custom error type usually has the suffix Error
type argError struct {
	arg int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	_, err := f(42)

	var ae *argError

	//errors.as is a more advanced version of errors.Is. Checks that a given error or any error in its chain matches a specific error type and converts it to a value of that type returning true
	//if no match reeturns false
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}