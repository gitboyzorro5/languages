package main

import "fmt"

func main(){

	var age int;

	fmt.Println("What is your age")

	fmt.Scan(&age)

	fmt.Printf("age plus 1 is %d",age + 1)
}