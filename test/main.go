package main

import "fmt"

type Animal struct {
	Name      string
	Real      bool
	Carnivore bool
}

type AquaticAnimal struct {
	Animal
	Fins int
}

type LandAnimal struct {
	Animal
	Legs int
}

func (a Animal) Output() {
	fmt.Printf("Name: '%s', Carnivore: %t, Real: %t", a.Name, a.Real, a.Carnivore)
}

func main() {
	cat := LandAnimal{}
	cat.Name = "Goon"
	cat.Real = true
	cat.Carnivore = true
	cat.Legs = 4
	cat.Output()
}

