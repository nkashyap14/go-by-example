package main

import "fmt"

//int parameter arguments are passed by value
func zeroval(ival int) {
	ival = 0
}

//pointer paprameter arguments are passed by reference have to dereference the pointer from the memory address
func zeroptr(iptr *int) {
	//derreference operator. used to access the value stored at memory address held by pointers
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	//passing in the address of i so basically passing in a pointer to int i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}