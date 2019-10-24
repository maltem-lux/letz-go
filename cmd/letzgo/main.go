package main

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/character"
)

func main() {

	// First way to create and instantiate a Character
	c1 := character.Character{}
	c1.Id = 1
	c1.Age = 33
	c1.Gender = "M"
	c1.Name = "Jeremy"

	// Second way to do it
	c2 := character.Character{Id:2, Age:29, Gender:"M", Name:"Jerôme"}

	fmt.Println(c1.ToString())
	fmt.Printf("%v", c2.ToString())

	// Or if you did not create a ToString method, you can either
	// Display all values of props
	fmt.Printf("\nOnly Property Values : %v", c1)
	// Or Display Prop/Value of the Char
	fmt.Printf("\nProperty Names + Values : %+v", c2)
}


