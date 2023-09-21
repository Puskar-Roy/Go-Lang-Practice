package main

import (
	"encoding/json"
	"fmt"
)

type user struct{
	Name      string
	Email     string
	Phone     int
	Password  string
	Cpassword string
}


func userJsonData()  {
	users := []user{
		{
			"Puskar Roy","puskarroy600@gmail.com",7449585365,"Puskar","Puskar",
		},
		{
			"Puskar Roy","puskarroy600@gmail.com",7449585365,"Puskar","Puskar",
		},
	}
	data,err := json.Marshal(users);
	if err!=nil{
		panic(err);
	}
	jsonData := string(data);
	fmt.Println(jsonData); 
}




func main() {
	userJsonData();

}