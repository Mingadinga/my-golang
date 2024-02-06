package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // 생략
	Tags     []string `json:"tags,omitempty"` // value가 null이면 생략
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"ANGULAR Bootcamp", 99, "LearnCodeOnline.in", "zxc123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// with indent
	finalJson, _ = json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
        "Price": 199,
		"website": "LearnCodeOnline.in",
		"tags": [ "full-stack", "js" ]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// just add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
