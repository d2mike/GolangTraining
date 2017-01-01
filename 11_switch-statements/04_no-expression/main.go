package main

import "fmt"

func main() {

	myFriendsName := "Al"

	switch {
	case len(myFriendsName) == 2: // Evaluates to true , and prints
		fmt.Println("wassap my friend with name of length 2")
	case myFriendsName == "Al": // Will not print this because order
		// matter, and the first one already evaluated to true
		fmt.Println("Wassap Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassap Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Jenny":
		fmt.Println("Wassap Jenny")
	case myFriendsName == "Sushant":
		fmt.Println("Wassap Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}
