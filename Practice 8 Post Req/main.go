package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func performPostReq() {
	const myurl string = "https://dummyjson.com/posts/add"

	requestBody := strings.NewReader(`{
			"userId":10
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	res , err := ioutil.ReadAll(response.Body);
	if err != nil {
		panic(err)
	}
	data:= string(res);
	fmt.Println(data);

}

func main() {
	fmt.Println("Hello Word")
	performPostReq()

}
