package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassap Tim")
	case "Jenny":
		fmt.Println("Wassap Jenny")
	case "Marcus":
		fmt.Println("Wassap Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassap Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassap Julian")
	case "Sushant":
		fmt.Println("Wassap Sushant")
	}
}
