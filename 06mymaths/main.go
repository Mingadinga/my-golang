package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4

	// 형변환
	sum := mynumberOne + int(mynumberTwo)
	fmt.Println("The sum is :", sum)

	// random number from math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random number from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
