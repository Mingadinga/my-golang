package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	// 생성
	languages := make(map[string]string) // [key]value

	// 추가
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// 참조
	fmt.Println("List of all languages :", languages) // map[JS:Javascript PY:Python RB:Ruby]
	fmt.Println("JS shorts for:", languages["JS"])    // Javascript

	// 삭제
	delete(languages, "RB")
	fmt.Println("List of all languages :", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// comma ok syntax
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

}
