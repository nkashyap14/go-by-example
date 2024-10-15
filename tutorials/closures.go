package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//nextInt becomes a returned function
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	//state of the returned function is unique to itself this value will be 1 the first time its called
	newInts := intSeq()
	fmt.Println(newInts())
}