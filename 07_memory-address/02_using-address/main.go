package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) // using "&" stores the input into a memory address
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, " yards.")
}
