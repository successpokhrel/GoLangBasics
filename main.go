package main

import "fmt"

func updateName(x string) {
	x = "dewge"
}

func updateNameTruePointer(x *string) {
	*x = "dewge"
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "fita"

	updateName(name)

	m := &name
	fmt.Println("memory address: ", m)
	fmt.Println("value at memory address: ", *m)

	updateNameTruePointer(m)

	fmt.Println("memory address: ", m)
	fmt.Println("value at memory address: ", *m)
}
