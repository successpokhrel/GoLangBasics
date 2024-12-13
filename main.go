package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["toffee pudding"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "biththal",
		546535453: "success",
		654646454: "Siddhu",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

}
