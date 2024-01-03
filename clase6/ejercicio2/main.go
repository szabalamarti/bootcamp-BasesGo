package main

import (
	"errors"
	"fmt"
)

const (
	Small  = "Small"
	Medium = "Medium"
	Large  = "Large"
)

type Product interface {
	Price() float64
}

type SmallProduct struct {
	Cost float64
}

func (s SmallProduct) Price() float64 {
	return s.Cost
}

type MediumProduct struct {
	Cost float64
}

func (m MediumProduct) Price() float64 {
	return m.Cost + 0.03*m.Cost
}

type LargeProduct struct {
	Cost float64
}

func (L LargeProduct) Price() float64 {
	return L.Cost + 0.06*L.Cost + 2500
}

func ProductFactory(productType string, cost float64) Product {
	switch productType {
	case Small:
		return SmallProduct{cost}
	case Medium:
		return MediumProduct{cost}
	case Large:
		return LargeProduct{cost}
	default:
		panic(errors.New("unssuported product type"))
	}
}

func main() {
	milk := ProductFactory(Small, 1000)
	chair := ProductFactory(Medium, 1500.0)
	car := ProductFactory(Large, 10000.0)
	fmt.Println("Milk (Small) Price:", milk.Price())
	fmt.Println("Chair (Medium) Price:", chair.Price())
	fmt.Println("Car (Large) Price:", car.Price())
}
