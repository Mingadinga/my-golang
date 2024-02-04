package main

import "fmt"

func main() {
	fmt.Println("Welcome to fuctions in golang")
	greeter()

	result := adder(3, 4)
	fmt.Println("Result is :", result)

	proResult, _ := proAdder(3, 4, 5, 6, 7)
	fmt.Println("Pro Result is :", proResult)
}

func greeter() {
	fmt.Println("Namstey from golang")
}

// 함수 내부에 함수 선언은 불가
func greeterTwo() {
	fmt.Println("Another method")
}

// 매개변수를 받고 반환값이 있는 함수
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// 가변인자를 받는 함수
// 두개 이상의 값을 반환하는 함수
func proAdder(values ...int) int, string {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Hi Pro result function"
}
