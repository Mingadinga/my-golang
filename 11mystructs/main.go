package main

import "fmt"

// 선언
type User struct {
	Name   string // public
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")
	// no inhreitance in golang!! : No super or parent

	// 생성
	hitesh := User{"Hitesth", "hitesh@go.dev", true, 16}

	// 출력 형식 : %+v
	fmt.Println(hitesh)                           // {Hitesth hitesh@go.dev true 16}
	fmt.Printf("hitesh details are: %+v", hitesh) // {Name:Hitesth Email:hitesh@go.dev Status:true Age:16}
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)

}
