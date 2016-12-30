package main

import "fmt"

const (
	A	= iota //0
	B	= iota //1
	C	= iota //2
)

//on another declaration the iota resets to zero
// do not have to re initalize iota
const (
	D	= iota //0
	E	       //1
	F	       //2
)

func main()  {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}