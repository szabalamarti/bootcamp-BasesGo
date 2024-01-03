package main

import "fmt"

type Student struct {
	Name     string
	Lastname string
	DNI      int
	Date     string
}

func (s Student) Detail() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Lastname:", s.Lastname)
	fmt.Println("DNI:", s.DNI)
	fmt.Println("Date:", s.Date)
}

func main() {
	student := Student{
		"Pedro", "Suarez", 11111, "30/01/1998",
	}
	student.Detail()
}
