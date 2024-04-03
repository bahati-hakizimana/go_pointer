package main

import "fmt"

func main() {

	fmt.Println("Well come to pointers Phase")

	// var ptr *int

	// fmt.Println("The value of this is", ptr)


	myNumber := 23

	var ptr = &myNumber


	fmt.Println("the actual value of pointer is", ptr)
	fmt.Println("the actual value of pointer is", *ptr)

	*ptr = *ptr * 4

	fmt.Println("The new value of pointer is :", myNumber)

}