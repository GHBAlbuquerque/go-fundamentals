package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

// implementing interfaces methods
func (d *Dog) Says() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	NumberOfTeeth int
}

// implementing interfaces methods
func (g *Gorilla) Says() string {
	return "Ugh!"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

func main() {
	dog := Dog{Name: "Toby", Breed: "German Shepherd"}
	gorilla := Gorilla{Name: "Mike", NumberOfTeeth: 38}

	PrintAboutAnimal(&dog)
	PrintAboutAnimal(&gorilla)
}

func PrintAboutAnimal(a Animal) {
	log.Println("Animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
