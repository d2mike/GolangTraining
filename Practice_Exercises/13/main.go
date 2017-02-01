package main

import "fmt"

func smallestDivisible() {
	var divisibleOneToTwenty bool
	var smallestDivisible int
	for divisibleOneToTwenty == false  {
		i := 0
		for i < 20 {
			if smallestDivisible % i == 0 {
				i++
			}
			if i == 20 {
				break
			} else {
				smallestDivisible++
			}
		}

	}
	fmt.Println("The smallest number divisble by 1- 20 is: ",
		smallestDivisible)
}

func main() {
	smallestDivisible()
}