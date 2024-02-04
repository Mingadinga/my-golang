package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	// 파일 생성, 존재한다면 덮어씀
	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	// 파일에 내용 작성
	io.WriteString(file, content)

	// 파일 읽기
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	// byte로 읽어옴!
	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) // shutdown
	}
}
