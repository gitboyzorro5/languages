package main 

import (

	"fmt"
	)

func main() {

	var x,y,i,j int;


	fmt.Println("Enter two integers in the range 1-20")

	fmt.Scanf("%d%d",&x,&y)

	for i = 1; i <= y; i++ {

      for j = 1; j <= x; j++ {

      fmt.Printf("@")

      }

     fmt.Printf("\n")
	}
}