package main

import (
	"testing"
	"os"
)

// test virus lethal
var vtest = MakeVirus("Ebola", 1.0, 1.0)

// test person id=0, unvaccinated
var ptest = MakePerson(false, 0, &vtest)

// lightweight test simulation
var populationTest = 100
var initialInfectedTest = 10
var vacPercentTest float32 = 0.5
var peopleTest = MakePeople()
var newlyInfectedTest []int
var currentInfectedTest []int

var f, err = os.Create(vtest.name + "_test")


var stest = MakeSimulation(
	f,
	peopleTest,
	newlyInfectedTest,
	vtest,
	populationTest,
	initialInfectedTest,
	currentInfectedTest,
	vacPercentTest,
)

func TestPopulate(t *testing.T) {
	stest.People = *Populate(stest.population, stest.initialInfected, stest.vacPercent, &stest.virus)
	if len(stest.People) != populationTest {
		t.Error("Test failed.  Populate does not populate the simulation with the correct number of people", len(peopleTest), populationTest)
	}	
}

func TestDidSurviveInfection(t *testing.T) {
	vtest.mortality = 1.0
	ptest.didSurviveInfection()
	if ptest.alive {
		t.Error("\n Test failed.  didSurviveInfection does not set person attributes correctly")
	}
}
func TestNumSurvivors(t *testing.T) {
	if stest.NumSurvivors() != 100 {
		t.Error("\n Test failed.  NumSurvivors does not reflect the right number of living people", stest.NumSurvivors(), 100)
	}
}

func TestShouldContinue(t *testing.T) {
	if !(stest.ShouldContinue()) {
		t.Error("\nTest failed. ShouldContinue evaluated to false when it should have returned true")
	}
}

func TestFindByID(t *testing.T) {
	i := stest.FindByID(0)
	if i.getID() != ptest.getID() {
		t.Error("\nTest failed", i.getID(), (ptest.getID()))
	}
}

func TestFindRandomPerson(t *testing.T) {
	if &ptest != stest.findRandomPerson() {
		t.Error("\nTest failed")
	}
}

func TestInfected(t *testing.T) {
	stest.infected(&ptest, &vtest)
	if ptest.alive {
		t.Error("\nTest failed.  infected did not infect target")
	}
	ptest.alive = true
	vtest.mortality = 0.0
	if !(ptest.alive) {
		t.Error("\nTest failed.  infected killed a person when it shouldn't have")
	}
}


