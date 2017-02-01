package main

import "fmt"

func main() {
	fmt.Print(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprintln(fname, lname), fmt.Sprintln(lname, fname)
}
