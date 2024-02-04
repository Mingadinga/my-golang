package main

import (
	"fmt"
	"net/http"
	"sync"
)

// 고루틴은 OS 스레드와 달리 고 런타임이 관리함. 더 가볍고 빠름
// 고루틴은 백그라운드에서 실행되므로, 고루틴을 기다려야만 고루틴이 실행된 것을 확인할 수 있음

// waiting group : 기다릴 고루틴 목록을 관리
// wg.Add(int) : 대기 그룹의 카운트 int만큼 추가
// defer wg.Done() : 대기 그룹의 카운트 1 감소
// wg.Wait() : 대기 그룹의 카운트가 0이 될 때까지 기다림

// mutext lock : 임계영역 접근 제어
// 고루틴 함수 내에서 임계영역을 mut.Lock(), mut.Unlock()으로 감싼다

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

var signals = []string{"test"}

func main() {

	// go 루틴 생성
	// go greeter("Hello") // 고루틴 생성, 백그라운드에서 실행됨
	// greeter("World")    // 함수가 종료되면서 프로그램이 종료되고, 아직 실행 중인 고루틴도 함께 종료되어 "Hello"를 출력하는 기회가 없어짐

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://gb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) // 대기 중인 고루틴의 수를 1 증가
	}

	// waiting group wait : 대기 그룹의 카운트가 0이 될 때까지 기다림
	wg.Wait() // usually end of the main method

	// 5개의 고루틴이 동시에 signals에 접근함
	// 고루틴은 고 런타임에 의해 관리되므로 락도
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		// 고루틴에서 잠시 재워서 main이 기다리도록 함
		// time.Sleep(2 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {

	// 고루틴 실행을 마치면 wg에 끝났다는 시그널을 보냄
	// 고루틴이 종료될 때 대기 그룹의 카운트를 1 감소
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint) // 임계 영역 : 여러 고루틴이 signals 메모리에 접근
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
