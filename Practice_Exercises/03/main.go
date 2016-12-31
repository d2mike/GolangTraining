package main

import "fmt"

func main()  {
	var firstname string
	var lastname string
	fmt.Print("Enter your name: ")
	fmt.Scan(&firstname, &lastname )

	fmt.Println("Hello,", firstname, lastname)
}