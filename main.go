package main

import "fmt"

func main() {

	//print formatting
	age := "12"
	name := "Thito"

	fmt.Printf("His name is %v and his age is %v\n", age, name)
	fmt.Printf("His name is %q and his age is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %f points!\n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("His name is %v and his age is %v\n", age, name)
	fmt.Println("saved: ", str)
}
