package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) //address in Hex

	var b *int = &a
	fmt.Println(b)  //address in hex
	fmt.Println(*b) //43

}

//b is an int pointer;
//b points to  the memory address where an int is stored
//to see the value in that memory address, add * in front of b
//this is dereferencing
//the * is an operator in this case
