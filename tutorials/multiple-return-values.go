package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//_ is the blank identifier
	_, c := vals()

	fmt.Println(c)
}