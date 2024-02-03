package main

import "fmt"

// 글로벌 영역에서는 var 생략 불가
// jwtToken := 300000

// public const
const LoginToken string = "ghsdfds"

func main() {
	// string
	var username string = "hwi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	// bool
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	// 1byte int
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	// 1byte int
	var smallFloat float64 = 255.343343
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// default valueds and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type : 암시적 선언
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style : 함수 내에서만 사용 가능
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	// use public const
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
