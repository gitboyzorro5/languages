package main 

import "fmt"

func main(){


	var num1 int;
	var num2 int;

	fmt.Println("Enter two numbers and i will tell yu the relationship they satisfy:")

	fmt.Scanf("%d %d", &num1, &num2)
	//fmt.Println("The two numbers are:",num1,num2)

	if (num1 == num2){

		fmt.Printf("%d is equal to %d ",num1,num2)

	}

	if (num1 != num2){

		fmt.Printf("%d is not equal to %d ",num1,num2)

	}

	if (num1 < num2){

		fmt.Printf("%d is less than %d ",num1,num2)

	}

	if (num1 > num2){

		fmt.Printf("%d is greater than %d ",num1,num2)

	}


}  