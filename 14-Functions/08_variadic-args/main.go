package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...) //dots are at end for variadic args, takes each
				// variable one at a time from a slice
	fmt.Println(n)
}

func average(sf ...float64) float64  {  // variadic dots are at front for
					// params, takes one value at a time.
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}