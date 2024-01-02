package main

import (
	"ejercicio5/internal/animals"
	"fmt"
)

func main() {
	animalDog, _ := animals.Animal("dog")
	animalCat, _ := animals.Animal("cat")
	animalHamster, _ := animals.Animal("hamster")

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(23)
	fmt.Println("Total de comid:", amount, "Kg")

	// Example of error
	// _, err := animals.Animal("anaconda")
	// if err != nil {
	// 	panic(err)
	// }
}
