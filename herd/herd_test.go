package herd 

import (
	"testing"
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

var stest = MakeSimulation(
	peopleTest,
	newlyInfectedTest,
	vtest,
	populationTest,
	initialInfectedTest,
	currentInfectedTest,
	vacPercentTest,
)

func TestPopulate(t *testing.T) {
	stest.Populate()
	if len(peopleTest) != populationTest {
		t.Error("Test failed.  Populate does not populate the simulation with the correct number of people", len(peopleTest), populationTest)
	}	
}

func TestDidSurviveInfection(t *testing.T) {
	ptest.didSurviveInfection()
	if ptest.alive {
		t.Error("Test failed.  didSurviveInfection does not set person attributes correctly")
	}
}
func TestNumSurvivors(t *testing.T) {
	if stest.NumSurvivors() != 100 {
		t.Error("Test failed.  NumSurvivors does not reflect the right number of living people", stest.NumSurvivors(), 100)
	}
}

func TestShouldContinue(t *testing.T) {
	if !(stest.ShouldContinue()) {
		t.Error("Test failed. ShouldContinue evaluated to false when it should have returned true")
	}
	stest.population = 0
	if stest.ShouldContinue() {
		t.Error("Test failed. ShouldContinue evaluated to true when it should have evaluated false")
	}
}

func TestFindByID(t *testing.T) {
	i := stest.FindByID(0)
	if i != &ptest {
		t.Error("Test failed")
	}
}

func TestFindRandomPerson(t *testing.T) {
	if &ptest != stest.findRandomPerson() {
		t.Error("Test failed")
	}
}

func TestInfected(t *testing.T) {
	stest.infected(&ptest, &vtest)
	if ptest.alive {
		t.Error("Test failed.  infected did not infect target")
	}
	ptest.alive = true
	vtest.mortality = 0.0
	if !(ptest.alive) {
		t.Error("Test failed.  infected killed a person when it shouldn't have")
	}
}
