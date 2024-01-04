package main

import (
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("customers.txt")

	defer func() {
		err = file.Close()
		fmt.Println("Ejecucion finalizada")
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	var line string
	for {
		_, err = fmt.Fscanln(file, &line)
		if err != nil {
			break
		}
		fmt.Println(line)
	}
}
