package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	// byte -> string
	//fmt.Println(string(content))

	// use strings.Builder : 재사용성 좋음
	var responseString strings.Builder
	content, _ = io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is :", byteCount)
	fmt.Println(responseString.String()) // byte -> string

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// create fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		}
	`)

	// post request with json data
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// read response from url
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// create fake form data
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		return
	}

	defer response.Body.Close()

	// check response
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
