package main

import (
	"encoding/json"
	"fmt"
)

 type users struct{
  Name      string
  Email     string
  Phone     int
  Password  string
  Cpassword string
}


func createJSON()  {
    jsonList := []users{
      { "Puskar Roy","puskar@gmail.com",7449585365,"Puskar","Puskar" },
    };

    byteJson,err := json.Marshal(jsonList);
    if err != nil{
      panic(err);
    }
    json := string(byteJson);

    fmt.Println(json)
    fmt.Printf("Type - %T",json)

}



func main() {
	createJSON();
}
