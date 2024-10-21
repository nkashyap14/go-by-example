package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("../defer.txt")
	//defer of the closing of the file with closeFile. will be executed at the end of the enclosing function (main) after write file has finished
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		os.Exit(1)
	}
}