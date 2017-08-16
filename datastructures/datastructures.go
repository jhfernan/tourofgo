package datastructures

import "fmt"

func RunSomeExamples() {
	var aValue float64 = 1.23
	var aPointer *float64 = &aValue
	fmt.Println("aPointer:", aPointer) // prints something like: "aPointer: 0xc42000e2e0"
	fmt.Println("*aPointer:", *aPointer) // prints "*aPointer: 1.23"
}
