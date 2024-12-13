package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "success", "biththal"}

	sort.Strings(names)
	fmt.Println(names)
}
