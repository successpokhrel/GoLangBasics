package main

import (
	"fmt"
)

var score = 99.5

func main() {
	sayHello("success")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

// go run main.go greetings.go
