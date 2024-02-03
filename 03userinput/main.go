package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	입출력 패키지 docs
	https://pkg.go.dev/bufio
*/

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n') // read until \n
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input) // string

}
