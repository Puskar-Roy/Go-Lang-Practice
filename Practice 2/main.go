package main

import "fmt"

func checkValue(a int) string {
	if a > 25 {
		return "Greater"
	} else {
		return "Smaller"
	}
}

func main() {
	var a int = 25
	var ans string = checkValue(a)
	fmt.Println(ans);
	fmt.Printf("Type - %T",ans)
}
