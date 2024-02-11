package main

import "fmt"

func main() {
	const A, B, C int = 1, 2, 4

	//type Gender int  constant cannot be changed by the program, but var can be changed
	//This data can only be of type boolean, number (integer, float, or complex) or string.
	const (
		UNNOWN = iota //assigne enumearatio 0, 1,2 ... to the values
		MALE
		FEMALE
		G
	)

	// if variable has to be exported, it must start with a capital letter
	number := 5 // initializing declaration inside function only

	//package scope is gloabal scope, declared outside of func
	//Variables of value type are contained in a stack memory.
	//values that are referenced are usually stored in the heap, which is garbage collected and which is much larger memory space than a stack
	fmt.Println(G, FEMALE, number)
}
