package main

import (
	"fmt"
	"os"
)

// Logger prints and logs program actions to a file
func Logger (f *os.File, s string) {
	fmt.Print(s)
	f.WriteString(s)
}

// Check checks for errors
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
