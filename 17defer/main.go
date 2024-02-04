package main

import "fmt"

func main() {
	// defer : 자신이 속한 함수의 실행이 끝난 후 호출됨
	// defer fmt.Println("World")
	// fmt.Println("Hello")

	// defer : 스택처럼 쌓을 수 있음
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	// 결과 : One Two Three

	// myDefer가 모두 끝나야 main의 defer가 호출됨
	myDefer()
}

func myDefer() {
	// myDefer가 종료될때 defer 실행
	// 0 1 2 3 4 -> 4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
