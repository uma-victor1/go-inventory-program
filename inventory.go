package main

import "fmt"

//import "fmt"

type bill struct {
	name string
	cart map[string]float64
	tip  float64
}

func createBill(name string) bill {
	b := bill{
		name: name,
		cart: map[string]float64{},
		tip:  0,
	}
	return b
}

func (b *bill) addTip(t float64) bill {
	b.tip = b.tip + t
	return *b
}

func (b *bill) addToCart(name string, price float64) {
	b.cart[name] = price
}

func (b bill) formatBill() string {
	fs := "Here is the bill breakdown for you \n"
	var total float64 = 0
	for i, p := range b.cart {
		b.cart["juice"] = 1.44
		if len(b.cart) > 0 {
			fs += fmt.Sprintf("%v...$%v \n", i, p)
			total += p
		}
	}
	//total
	fs += fmt.Sprintf("Total is %v...", total)
	return fs
}
