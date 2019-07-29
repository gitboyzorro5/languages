package main

import "fmt"

func main(){
	
 color := "violet"

 switch color{
 	case "red":
 	fmt.Printf("color is red")
 	case "blue":
 	fmt.Printf("color is blue")
 	default:
 	fmt.Printf("color is not red or blue")
 }
}