package main

import "fmt"

var x int = 42 //scope of "x" is whole package because it is not in func main()
//accessible throughout package

func main() {
	fmt.Println(x)
	foo()
}
func foo() {
	fmt.Println(x)
}
