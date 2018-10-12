package main

import (
	"fmt"
	"strconv"
	"./herd"
)

func main() {
	// CLI args
	// virusName := os.Args[3]

	// validate and convert
	// population := I(os.Args[1])
	// vacPercent := F32(os.Args[2])
	// initialInfected := I(os.Args[6])
	// mortalityRate := F32(os.Args[4])
	// basicReproNum := F32(os.Args[5])

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