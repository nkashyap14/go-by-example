package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	fmt.Printf("struct1: %v\n", p)

	//+v variant includes the struct field names if the value is a struct
	fmt.Printf("struct2: %+v\n", p)

	//prints the source code snipppet that owuld produce that value
	fmt.Printf("struct3: %#v\n", p)

	//to print type
	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	//%d for standard base-10 formatting
	fmt.Printf("int: %d\n", 123)

	//prints a binary representation
	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)
	
	//for hex encoding
	fmt.Printf("hex: %x\n", 456)

	//basic float formatting
	fmt.Printf("float1: %f\n", 78.9)

	//slightly different versions of scientific notation formatting
	fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)

	//basic string printing
    fmt.Printf("str1: %s\n", "\"string\"")

	//to doublequote a string use %q
    fmt.Printf("str2: %q\n", "\"string\"")

	//%x converts even a string to hex
    fmt.Printf("str3: %x\n", "hex this")

	//printing representation of a pointer
    fmt.Printf("pointer: %p\n", &p)

	//use number after % in the verb to specify the width of an integer
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	//for floats can use width.precision syntax
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	//- used to left justify
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	//can use width after % even with strings
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	//similar can use lest justification
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	//Sprintf formats and returns a string without printing it anywhere
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

	//can write to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}