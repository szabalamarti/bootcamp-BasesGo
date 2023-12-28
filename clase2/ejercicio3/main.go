package main

import "fmt"

// Se puede usar un array. Otras formas son un switch con cada numero o muchos if/else
var months []string = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func main() {
	fmt.Println("Ingresa el numero del mes: ")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error de lectura")
		return
	}
	if i < 1 || i > 12 {
		fmt.Println("El mes esta fuera del rango")
		return
	}
	month_name := GetMonthName(i)
	fmt.Println("El mes es:", month_name)
}

func GetMonthName(month int) (month_name string) {
	return months[month]
}
