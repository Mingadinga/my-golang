package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// 선언과 함께 초기화
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList) // []string

	// 요소 추가 : [Apple Tomato Peach Mango Banana]
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// 슬라이싱 : [Tomato Peach Mango Banana]
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// 슬라이싱 : [Peach Mango]
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// make로 생성
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 123
	highScores[3] = 456

	// highScores[4] = 777 // out of range
	highScores = append(highScores, 555, 666) // 가변적으로 추가 가능

	// 값 자체를 sort
	sort.Ints(highScores)
	fmt.Println(highScores)                     // [123 234 345 456 555 666]
	fmt.Println(sort.IntsAreSorted(highScores)) // true

	// how to remove a value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
