package main

import (
	"fmt"
)

func main() {
	var controller string
	var num1,num2 int
	fmt.Println("Enter First Number ")
	_, err1 := fmt.Scanf("%d\n", &num1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("Enter Second Number")
	_, err2 := fmt.Scanf("%d\n", &num2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Enter Your Sign")
	_, err3 := fmt.Scanf("%s\n", &controller)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	if(controller == "+"){
		ans := num1+num2;
		fmt.Println("Ans Is ",ans);
	}else if(controller == "-"){
		ans := num1-num2;
		fmt.Println("Ans Is = ",ans);
	}

}
