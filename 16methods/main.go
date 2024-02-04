package main

import "fmt"

func main() {
	hitesh := User{"Hitesth", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitest details are %+v\n", hitesh)
	fmt.Printf("hitest is %v and email is %v\n", hitesh.Name, hitesh.Email)

	// 메소드처럼 호출!
	hitesh.GetStatus() // 읽기
	hitesh.NewMail()   // 쓰기

	// 쓰기 작업을 하는 메소드를 호출해도 실제 값은 변하지 않음. 값이 복사되어 넘어감
	fmt.Printf("hitest is %v and email is %v\n", hitesh.Name, hitesh.Email)

	// 실제로 바꾸려면 포인터를 넘겨야함
	hitesh.NewMailByPointer()
	fmt.Printf("hitest is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// func와 이름 사이에 참조를 추가하면 메소드처럼 호출 가능
func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}

// type에 메소드를 연결하는 컨셉인 듯, 값 복사
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}

// 실제 값을 변경하려면 포인터 사용 필요
func (u *User) NewMailByPointer() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}
