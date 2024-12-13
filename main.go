package main

import "fmt"

func main() {
	mybill := newBill("success's bill")

	fmt.Println(mybill.format())
}

// go run main.go bill.go
