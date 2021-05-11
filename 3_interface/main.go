package main

import (
	"fmt"
	"strconv"
)

// Interfaces define behaviours that are expected and does not care about the actual values or types

type Announcer interface {
	MakeSound()
}

type LegsCounter interface {
	CountLegs() int
}

// Interfaces can be combined
type Animal interface {
	Announcer
	LegsCounter
}

// Cat implements all 3 interfaces because it has both the MakeSound and CountLegs method
type Cat struct {}

func (cat Cat) MakeSound() {
	fmt.Println("Meow!")
}

func (cat Cat) CountLegs() int {
	return 4
}

// Car only implements the Announcer interface
type Car struct {
	Sound string
}

func (car Car) MakeSound() {
	fmt.Println(car.Sound)
}

func main()  {
	// Side note - functions can be set to variables as well as values in a struct
	announcerFunc := func(announcer Announcer) {
		announcer.MakeSound()
	}

	animalFunc := func(animal Animal) {
		for i := 0; i < 3; i++ {
			animal.MakeSound()
		}
		fmt.Println("this animal has " + strconv.Itoa(animal.CountLegs()) + " legs!")
	}

	cat := Cat{}
	car := Car{
		Sound: "Beep!",
	}

	announcerFunc(cat)
	announcerFunc(car)

	animalFunc(cat)
	// Car does not implement the LegCounter interface therefore it cannot be used as an Animal
	//animalFunc(car)
}