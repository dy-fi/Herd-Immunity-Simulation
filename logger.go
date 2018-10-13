package main

import (
	"fmt"
)

// Logger prints and logs program actions to a file
func (sim Simulation) Logger (s string) {
	fmt.Print(s)
	sim.f.WriteString(s)
}

// Check checks for errors
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
