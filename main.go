package main

import (
	"fmt"
	"strconv"
	"os"
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

	virusName := "EbolAIDS"
	var mortalityRate float32 = 0.2
	var basicReproNum float32 = 0.6

	// set file
	var f, err = os.Create(virusName + "_log")
	Check(err)
	defer f.Close()

	population := 10000
	initialInfected := 1000
	var vacPercent float32 = 0.6

	// create virus
	virus := MakeVirus(
		virusName,
		mortalityRate,
		basicReproNum,
	)

	// create simulation
	var people = MakePeople()
	var newlyInfected []int
	var currentInfected []int

	sim := MakeSimulation(
		f,
		people,
		newlyInfected,
		virus,
		population,
		initialInfected,
		currentInfected,
		vacPercent,
	)

	// run loop
	for sim.ShouldContinue() == true {
		sim.Timestep()
		fmt.Printf("Timestep  |  People alive: %d", sim.NumSurvivors())
		Logger(sim.f, "Survivors: " + strconv.Itoa(sim.NumSurvivors()))
	}
}