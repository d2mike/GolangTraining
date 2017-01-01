package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassap Daniel")
	case "Medhi":
		fmt.Println("Wassap Medhi")
	case "Jenny":
		fmt.Println("Wassap Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
no default fallthrough
fallthrough is optional
-- you can specify fallthrough by explicitly stating it
-- break isn't needed like in other languages
*/
