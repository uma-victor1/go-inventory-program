package main

//import "fmt"


type bill struct {
	name string
	cart map[string]float64
	tip float64
}

func createBill(name string) bill {
	b := bill{
		name: name,
		cart: map[string]float64{},
		tip: 0,
	}
	return b
}