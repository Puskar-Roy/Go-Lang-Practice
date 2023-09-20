package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const myurl string = "https://dummyjson.com/products/1"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// qparams := result.Query();
	// fmt.Println(qparams)
	request, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(request)
	defer request.Body.Close()
	response, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("This Is The Encoded Response ", response)
	res := string(response)
	fmt.Println(res)

}
