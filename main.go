package main

import (
	"fmt"
	"strconv"

	"./herd"
)

func main() {
	// // CLI args
	// virusName := os.Args[3]

	// // validate and convert
	// population := validateInt(os.Args[1])
	// vacPercent := validateFloat32(os.Args[2])
	// initialInfected := validateInt(os.Args[6])
	// mortalityRate := validateFloat32(os.Args[4])
	// basicReproNum := validateFloat32(os.Args[5])

	herd.Logger()

	virusName := "EbolAIDS"
	var mortalityRate float32 = 0.2
	var basicReproNum float32 = 0.6

	population := 10000
	initialInfected := 1000
	var vacPercent float32 = 0.6

	// create virus
	virus := herd.MakeVirus(
		virusName,
		mortalityRate,
		basicReproNum,
	)

	// create simulation
	people := herd.MakePeople()
	var newlyInfected []int
	var currentInfected []int

	sim := herd.MakeSimulation(
		people,
		newlyInfected,
		virus,
		population,
		initialInfected,
		currentInfected,
		vacPercent,
	)


	sim.People = sim.Populate()
	
	// run loop
	for sim.ShouldContinue() {
		sim.Timestep()
		fmt.Printf("Timestep  |  People alive: %d", sim.NumSurvivors())
		herd.Log <- "Survivors: " + strconv.Itoa(sim.NumSurvivors())
	}
}

func validateInt(s string) int {
	if word, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return word
	}
}

func validateFloat32(f string) float32 {
	if word, err := strconv.ParseFloat(f, 32); err != nil {
		panic(err)
	} else {
		return float32(word)
	}
}
