package main

import "fmt"

func main() {
	//fmt.Println("hi")
	const A, B, C int = 1, 2, 4

	//type Gender int  constant cannot be changed by the program
	//This data can only be of type boolean, number (integer, float, or complex) or string.
	const (
		UNNOWN = iota
		MALE
		FEMALE
		G
	)

	// if variable has to be exported, it must start with a capital letter

	fmt.Println(G, FEMALE)
}
