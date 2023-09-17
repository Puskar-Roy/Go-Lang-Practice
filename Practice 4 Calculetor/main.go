package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func subtract(num1 int, num2 int) int {
	return num1 - num2
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func division(num1 int, num2 int) int {
	return num1 / num2
}

func main() {
	var controller string
	var num1, num2 int

	fmt.Println("Enter First ")
	_, err := fmt.Scanf("%d\n", &num1)
	if( err != nil ){
		fmt.Println(err)
		return
	}
	fmt.Println("Enter Second ")
	_, err1 := fmt.Scanf("%d\n", &num2)
	if( err1 != nil ){
		fmt.Println(err1)
		return
	}
	fmt.Println("Enter Controller ")
	_, err2 := fmt.Scanf("%s\n", &controller)
	if( err2 != nil ){
		fmt.Println(err2)
		return
	}

	if(controller == "+"){
		ans := add(num1,num2);
		fmt.Println("Ans = ",ans)
	}else if(controller == "-"){
		ans := subtract(num1,num2);
		fmt.Println("Ans = ",ans)
	}else if(controller == "*"){
		ans := multiply(num1,num2);
		fmt.Println("Ans = ",ans)
	}else{
		ans := division(num1,num2);
		fmt.Println("Ans = ",ans)
	}

	
}
