package main

import "fmt"

func main() {
	mybill := newBill("success's bill")

	mybill.addItem("Pancake", 5.55)
	mybill.addItem("Sel roti", 2.75)
	mybill.addItem("Samosa", 3.15)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}

// go run main.go bill.go
