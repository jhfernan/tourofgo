package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	a := "Hello, world!"

	fmt.Println(a)
	fmt.Println(sayHello("Jonathan"))

	// setting square root function variables
	squareRoot, err := squareRoot(9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(squareRoot)
}

// sample functions
func sayHello(name string) (greeting string) {
	greeting = "Hello, " + name + "!"
	return
}

// returning multiple values
func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("Can't have a square root of a negative number")
	}
	return math.Sqrt(x), nil
}
