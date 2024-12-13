package main

import (
	"fmt"
	"math"
)

func sayNamaste(n string) {
	fmt.Printf("Namaste %v !\n", n)
}

func sayFeriBhetaula(n string) {
	fmt.Printf("Feri bhetaula %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayNamaste("sano keto")
	// sayNamaste("Thule")
	// sayFeriBhetaula("sano keto")

	cycleNames([]string{"Biththal", "Ronit d", "Nikhil d"}, sayNamaste)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Printf("%0.2f, %0.2f \n", a1, a2)
}
