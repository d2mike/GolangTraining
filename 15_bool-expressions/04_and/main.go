package main

import "fmt"

func main() {

	if true && false {
		fmt.Println("This did not run")
	} else {
		println("This did run")
	}
}
