package main

import (
	"fmt"
)

func main() {
	for i:= 1; i <= 5; i++ {
		fmt.Println(i)
	}

	if false {
		fmt.Println("true")
	} else if true {
		fmt.Println("Second true")
	} else {
		fmt.Println("else true")
	}
}
