package main

import(
	"math/rand"
	"strconv"
	"os"
)

// Virus struct
type Virus struct {
	name string
	mortality float32
	repro float32
}

// Simulation struct
type Simulation struct {
	f *os.File
	People []Person 
	newlyInfected []int 
	virus Virus
	population int

	// constructor assignments
	initialInfected int
	currentInfected []int 
	vacPercent float32
}

// Person struct 
type Person struct {
	vac bool
	id int

	// constructor assignments
	alive bool
 	virus *Virus
}

// Constructors

// MakeVirus virus constructor
func MakeVirus(name string, mortality float32, repro float32) Virus {
	v := Virus{name, mortality, repro}
	return v
}

//MakeSimulation simulation constructor
func MakeSimulation(f *os.File, people []Person, newlyInfected []int, virus Virus, population int, initialInfected int, currentInfected []int, vacPercent float32) Simulation {
	ppl := Populate(population, initialInfected, vacPercent, &virus)
	s := Simulation{f, *ppl, newlyInfected, virus, population, initialInfected, currentInfected, vacPercent }
	return s
}

// MakePerson person constructor
func MakePerson(mVac bool, mID int, mInf *Virus) Person {
	p := Person{mVac, mID, true, mInf}
	return p
}

// MakePeople people initializer 
func MakePeople() []Person {
	var ppl []Person 
	return ppl 
}

// Converters

// S converts an int to a string
func S(i int) string {
	s := strconv.Itoa(i)
	return s
}

// I converts a string to an int
func I(s string) int {
	if word, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return word
	}
}

// F32 converts a string to a float32
func F32(f string) float32 {
	if word, err := strconv.ParseFloat(f, 32); err != nil {
		panic(err)
	} else {
		return float32(word)
	}
}

// Person methods

// getter
func (p Person) getID () int {
	return p.id
}

// InfectionSurvivalChance represents the survival chance
func (p Person) didSurviveInfection() (bool) {
	if p.virus.mortality > rand.Float32() {
		p.alive = false
		return false
	}
	return true
}


// Simulation methods

// NumSurvivors return the number of people alive in the simulation
func (sim Simulation) NumSurvivors() int {
	survivors := 0
	for _,p := range(sim.People) {
		if p.alive {
			survivors++ 
		}
	}
	return survivors 
}

// ShouldContinue checks if program should continue
func (sim Simulation) ShouldContinue() bool {
	if sim.NumSurvivors() > 0 {
		Logger(sim.f,  "ShouldContinue: True")
		return true
	}
	Logger(sim.f,  "Simulation ending...")
	// Go has automatic garbage collection so ending without deleting objects is okay
	return false
}

// Populate returns a list of ints to be used as the people list 
func Populate(pop int, inf int, vac float32, v *Virus) *[]Person {
	nextID := 0
	var pSlice = make([]Person, pop)

	vacd := int(float32(pop) * vac) 
	remaining := pop - (vacd + inf)

	for i := 0; i < vacd; i++ {
		pSlice = append(pSlice, MakePerson(true, nextID, nil))
	} 
	for i := 0; i < inf; i++ {
		pSlice = append(pSlice, MakePerson(true, nextID, v))
	}
	for i := 0; i < remaining; i++ {
		pSlice = append(pSlice, MakePerson(false, nextID, nil))
	}
	return &pSlice
}

// FindByID finds a user by their id
func (sim Simulation) FindByID(id int) *Person {
	for _,p := range sim.People {
		if p.id == id {
			return &p 
		}
	}
	return nil
}

// returns a random person
func (sim Simulation) findRandomPerson() *Person {
	return sim.FindByID(rand.Intn(len(sim.People)))
}

// died of disease chance
func (sim Simulation) infected(per *Person, vir *Virus) {
	// survival chance
	if rand.Float32() >= sim.virus.repro {
		per.alive = false
		Logger(sim.f,  S(per.id) + " died from infection")
	}
	// appends the id to the newly infected index
	sim.newlyInfected = append(sim.newlyInfected, per.id)
	Logger(sim.f,  S(per.id) + " became a host")
}

// interaction between an infected and healthy non-vacced person
func (sim Simulation) interact(pArg1, pArg2 int) {
	p1 := sim.FindByID(pArg1)
	p2 := sim.FindByID(pArg2)
	// check both are alive
	if  p1.alive && p2.alive {
		// if p2 is vaccinated or has the virus do nothing
		if p2.vac || p2.virus != nil {
			Logger(sim.f,  "Interaction between " + S(p1.id) + " and " + S(p2.id) + " uneventful")
			return
		}
		// else infect p2
		sim.infected(p2, p1.virus)
		Logger(sim.f,  S(p1.id) + " infected " + S(p1.id))
	}
} 

// Timestep represents 1 exposure period
func (sim Simulation) Timestep () {
	Logger(sim.f,  "Timestepping...")
	for _,p := range sim.currentInfected {
		for i := 100; i > 0; i-- {
			sim.interact(p, sim.findRandomPerson().getID())
		}
	}
	// dump newly infected people into current infected list
	for _,i := range sim.newlyInfected {
		if (sim.FindByID(i).didSurviveInfection()) {
			sim.currentInfected = append(sim.currentInfected, sim.FindByID(i).getID())
			Logger(sim.f,  "newly infected added to current infected list")
		}
	}
}
