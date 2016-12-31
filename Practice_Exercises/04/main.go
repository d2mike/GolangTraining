package main

import "fmt"

func main()  {
	var smallnumba int
	var largenumba int
	fmt.Print("Choose a small number: ")
	fmt.Scan(&smallnumba)
	fmt.Print("Choose a large number: ")
	fmt.Scan(&largenumba)

	remainder := largenumba % smallnumba
	println("The remainder of", largenumba, "divided by", smallnumba, "is:", remainder)

}
