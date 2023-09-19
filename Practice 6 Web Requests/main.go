package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Reqs");
	response ,err := http.Get(url);
	if err!=nil{
		panic(err);
	}
	
	fmt.Printf("Response Type Is %T",response);
	defer response.Body.Close();
	data , err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err);
	}
	fmt.Println("Response Is = ",string(data))
}