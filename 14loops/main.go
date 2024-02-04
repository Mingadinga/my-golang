package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	// slice loop
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// index loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// index loop with range
	for i := range days {
		fmt.Println(days[i])
	}

	// index and value loop (자주 사용!)
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// comma ok
	for _, day := range days {
		fmt.Printf("index is and value is %v\n", day)
	}

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	// goto label
lco:
	fmt.Println("Jumping at LearnCodeonLine.in")

}
