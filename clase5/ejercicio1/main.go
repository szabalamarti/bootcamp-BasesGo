package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product = []Product{
	{
		ID:          1,
		Name:        "A",
		Price:       1.0,
		Description: "Food A",
		Category:    "Food",
	},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, p := range Products {
		fmt.Println(p)
	}
}

func getById(id int) Product {
	for _, p := range Products {
		if p.ID == id {
			return p
		}
	}
	return Product{}
}

func main() {
	pizza := Product{
		ID:          2,
		Name:        "Pizza",
		Price:       5.0,
		Description: "Nice Pizza",
		Category:    "Food",
	}
	pizza.Save()
	pizza.GetAll()
	fmt.Println(getById(1))
}
