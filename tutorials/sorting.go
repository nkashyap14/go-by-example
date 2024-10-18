package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "a", "b"}

	//sorting functions work for anny ordered built in type
	slices.Sort(strs)

	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	//can use slices package to check if a slice is in sorted order
	s := slices.IsSorted(ints)
	fmt.Println("Sorted:   ", s)
}