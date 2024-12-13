package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("X is: ", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is: ", i)
	// }

	names := []string{"mario", "lungi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		//this doesnot update the value in original slice.
		value = "new string"
	}

	fmt.Println(names)

}
