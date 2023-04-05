package main

import "fmt"

func main() {
	fmt.Println("Valid go program here")
	myBill := createBill("Uma Victor")
	myBill.addToCart("Juice", 9.44)
	myBill.addToCart("orange", 5.99)
	myBill.addTip(3.5)
	fmt.Println(myBill)
}
