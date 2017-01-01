package main

import "fmt"

func main() {
	i := 0 //initlize
	for {
		fmt.Println(i)
		if i >= 10 { //condition
			break //will stop when it reaches 10
		}
		i++ //post
	}
}
