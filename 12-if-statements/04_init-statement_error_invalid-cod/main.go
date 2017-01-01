package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	fmt.Println(food) // initialization is to keep the scope tight, so
	// we can see here that we can't get the variable "food" out of the
	// code block
}
