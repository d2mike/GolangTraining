package main

import "fmt"

func big (numbers...int) int {
	var largest int

	for _, v := range numbers{
		if v > largest{
			largest = v
		}
	}
	return largest
}

func main()  {
	greatest := big(154)
	fmt.Println(greatest)
}