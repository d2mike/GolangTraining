package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++ //will run forever until someone tells it to stop
	}
}
