package main

import "fmt"

func modelTraing()  {
	var dataset1 , dataset2 , dataset3 , temp int;
	fmt.Print("Enter DataSet - ");
	fmt.Scanf("%d\n",&dataset1);
	fmt.Print("Enter DataSet - ");
	fmt.Scanf("%d\n",&dataset2);
	fmt.Print("Enter DataSet - ");
	fmt.Scanf("%d\n",&dataset3);
	fmt.Println("Actual Numbers = ",dataset1," ",dataset2," ",dataset3)
	temp = dataset1;
	dataset1 = dataset2;
	dataset2 = temp;
	fmt.Println("\t After Swaping ")
	fmt.Println("Actual Numbers = ",dataset1," ",dataset2," ",dataset3)

}


func main() {
	modelTraing();
}