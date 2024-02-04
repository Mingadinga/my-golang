package main

import (
	"fmt"
	"sync"
)

// 경쟁조건 제어를 위해 매번 뮤텍스 락을 사용하는 것은 수고로움
// 채널 : 여러 고루틴 사이에 데이터를 주고 받을 때 사용. 각 고루틴은 독립적으로 동작함
// 우와 이거 정말 특이하다..!!!
// 메시지 채널을 둬서 경쟁조건을 해결한다는 아이디어는 싱글 스레드 기반의 레디스로 경쟁조건을 없애는 논리와 비슷한 듯
// 메시지 채널에서 pub sub 주체가 있고 여러 채널이 생성되고 닫힐 수 있다는 점에서 카프카와 비슷하다고 느낌

func main() {
	fmt.Println("Channels in golang - LearnCodeOnline.in")

	// int 채널 생성
	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	// 채널에 데이터 5를 보냄
	// fatal error: all goroutines are asleep - deadlock!
	// 아무도 소비하지 않아서 메인 고루틴이 대기하게 됨, 데드락 발생
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	// read ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh // 채널의 데이터 소비
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)

	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5 // 채널에 데이터를 보냄
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
