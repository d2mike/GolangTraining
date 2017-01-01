package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \t", a) //prints value in default format
	fmt.Printf("%T \n", a) //prints value type ie: String, int, bool
	fmt.Printf("%v \t", b)
	fmt.Printf("%T \n", b)
	fmt.Printf("%v \t", c)
	fmt.Printf("%T \n", c)
	fmt.Printf("%v \t", d)
	fmt.Printf("%T \n", d)
}
