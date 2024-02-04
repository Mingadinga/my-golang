package main

import (
	"fmt"
	"sync"
)

// race conditon 확인 : go run --race .
// race conditon : 두 고루틴이 동시에 현재 score의 길이를 읽어오고, 이를 기반으로 새로운 배열을 만들어 값을 추가
// 해결 : 임계영역에 mutex 추가
func main() {
	fmt.Println("Race condition - LearnCodeonline.in")

	waitgroup := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	// 익명함수 고루틴 실행
	waitgroup.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")

		m.Lock()
		score = append(score, 1)
		m.Unlock()

		wg.Done()
	}(waitgroup, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")

		m.Lock()
		score = append(score, 2)
		m.Unlock()

		wg.Done()
	}(waitgroup, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")

		m.Lock()
		score = append(score, 3)
		m.Unlock()

		wg.Done()
	}(waitgroup, mut)

	waitgroup.Wait()

	fmt.Println(score)

}
