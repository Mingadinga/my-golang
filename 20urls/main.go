package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=ghbj456

	// query parameters
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) // url.Values
	fmt.Println(qparams)                                      // map[coursename:[reactjs] paymentid:[ghbj456]]
	fmt.Println(qparams["coursename"])                        // [reactjs]

	for key, val := range qparams {
		fmt.Println("Key is :", key, "Param is :", val)
	}
	// Key is : coursename Param is : [reactjs]
	// Key is : paymentid Param is : [ghbj456]

	// URL 생성
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL) // https://lco.dev/tutcss

}
