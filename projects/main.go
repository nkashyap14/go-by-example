package main

import (
	"consistent-hashing/hashing"
	"fmt"
)

func main() {
    ring := hashing.NewRing(5, 3, nil)
	fmt.Println(ring)
}