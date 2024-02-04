package main

import "fmt"

// 의외로 배열은 go에서 잘 사용되지 않음
func main() {
	fmt.Println("Welcome to array in golangs")
	var fruitList [4]string // fixed

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("vegy list is:", vegList)
	fmt.Println("vegy list is:", len(vegList))

}
