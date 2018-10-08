package herd

import(
	"math/rand"
)

// Virus struct
type Virus struct {
	name string
	mortality float32
	repro float32
}

// Simulation struct
type Simulation struct {
	people []int
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
func MakeSimulation(people []int, newlyInfected []int, virus Virus, population int, initialInfected int, currentInfected []int, vacPercent float32) Simulation {
	s := Simulation{people, newlyInfected, virus, population, initialInfected, currentInfected, vacPercent }
	return s
}

// MakePerson person constructor
func MakePerson(mVac bool, mID int, mInf *Virus) Person {
	p := Person{mVac, mID, true, mInf}
	return p
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
	for _,p := range(sim.people) {
		if sim.FindByID(p).alive {
			survivors++ 
		}
	}
	return survivors 
}

// ShouldContinue checks if program should continue
func (sim Simulation) ShouldContinue() bool {
	if sim.NumSurvivors() > 0 {
		return true 
	}
	return false
}

// Populate returns a list of ints to be used as the people list 
func (sim Simulation) Populate() {
	nextID := 0
	v := &sim.virus
	vacd := int(float32(sim.population) * sim.vacPercent)
		
	// populate vaccinated individuals
	for i := 0; i < vacd; i++ {
		sim.people = append(sim.people, MakePerson(true, nextID, nil).id)
		nextID++ 
	}
	// populate initially infected
	for i := 0; i < sim.initialInfected; i++ {
		sim.people = append(sim.people, MakePerson(false, nextID, v).id )
		nextID++
	}
}

// FindByID finds a user by their id
func (sim Simulation) FindByID(id int) *Person {
	for _,p := range sim.people {
		if p == id {
			return sim.FindByID(p)
		}
	}
	return nil
}

// returns a random person
func (sim Simulation) findRandomPerson() *Person {
	return sim.FindByID(rand.Intn(100))
}

// died of disease chance
func (sim Simulation) infected(per *Person, vir *Virus) {
	// survival chance
	if rand.Float32() >= sim.virus.repro {
		per.alive = false
	}
	// appends the id to the newly infected index
	sim.newlyInfected = append(sim.newlyInfected, per.id)
}

// interaction between an infected and healthy non-vacced person
func (sim Simulation) interact(pArg1, pArg2 int) {
	p1 := sim.FindByID(pArg1)
	p2 := sim.FindByID(pArg2)
	// check both are alive
	if  p1.alive && p2.alive {
		// if p2 is vaccinated or has the virus do nothing
		if p2.vac || p2.virus != nil {
			return 
		}
		// else infect p2
		sim.infected(p2, p1.virus)
	}
} 

// Timestep represents 1 exposure period
func (sim Simulation) Timestep () {
	for _,p := range sim.currentInfected {
		for i := 100; i > 0; i-- {
			sim.interact(p, sim.findRandomPerson().getID())
		}
	}
	// dump newly infected people into current infected list
	for _,i := range sim.newlyInfected {
		if(sim.FindByID(i).didSurviveInfection()){
			sim.currentInfected = append(sim.currentInfected, sim.FindByID(i).getID())
		}
	}
}