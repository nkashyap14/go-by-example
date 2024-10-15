package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//go treats strings as containers of text encoded in UTF-8 while in other languages strings are made up of characters
	// in go a character is called a rune. An integer that represents a Unicode code point

	const s = "สวัสดี"

	//will produce the lenght of the raw bytes stored within as strings are equivalent to []byte
	fmt.Println("Len:", len(s))


	//generates hex values of all the bytes that constitute code points in s
	for i := 0; i < len(s); i ++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	//to count runes we need to use utf8 package
	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	//can achieve same as above by using utf8.DecodeRuneInString
	fmt.Println("Using utf8.DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		//w represents the width of the rune in terms of code points. we have to to increment i by that
		w = width

		//passing a rune value to a function
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	//value enclosed in single quotes are rune literals. can compare a rune value to a rune literal directly
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
        fmt.Println("found so sua")
    } 
}