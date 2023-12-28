package main

import "fmt"

func main() {
	var word string
	fmt.Println("Ingresa tu palabra: ")
	fmt.Scanln(&word)

	// Count word
	fmt.Println("Largo:", len(word))

	// Spell word
	for i := 0; i < len(word); i++ {
		fmt.Printf("%s ", string(word[i]))
	}
	fmt.Println("")
}
