package controlflow

import (
	"fmt"
)

func RunSomeExamples() {
	// for loop
	for i:= 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// if else statements
	if false {
		fmt.Println("true")
	} else if true {
		fmt.Println("Second true")
	} else {
		fmt.Println("else true")
	}

	// sample switch statement
	houseNumber := 2404

	switch houseNumber {
	case 2404:
		fmt.Println("That's our house!")
	case 2424:
		fmt.Println("Our Neighbor's house!")
	default:
		fmt.Println("I don't know whose house that is...")
	}
}
