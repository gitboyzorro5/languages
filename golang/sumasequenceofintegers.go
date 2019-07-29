package main 

import ("fmt")

func  main() {
	
var counter int;

var sum int = 0;

var holder int;


fmt.Println("Enter the number of values to be summed:")

fmt.Scanf("%d",&counter)

for i :=1 ; i <= counter; i++{

	fmt.Println("Enter a number")

	fmt.Scanf("%d",&holder)

	sum = sum + holder


} 

fmt.Printf("The sum is %d",sum)

}