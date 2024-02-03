package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("Value of pointer is", ptr) // nil

	myNumber := 23
	var ptr2 *int = &myNumber
	fmt.Println("Value of actual pointer is", ptr2)  // 0x1400010a008
	fmt.Println("Value of actual pointer is", *ptr2) // 23

	// left : 주소 참조
	// right : 값 참조
	*ptr2 = *ptr2 * 2
	fmt.Println("New value is:", myNumber) // 46

}
