package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of go lang")

	// 현재 시간
	presentTime := time.Now()
	fmt.Println(presentTime)

	// date format의 필드는 항상 고정
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// year, month, day, hour, minute, second, nano second
	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
