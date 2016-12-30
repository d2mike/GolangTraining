package main

import "fmt"

func main()  {
	x := 42
	fmt.Println(x) //scope of x is limited to func main() no further than closing Brace
	foo()
}

func foo()  {
	fmt.Println(x)
}