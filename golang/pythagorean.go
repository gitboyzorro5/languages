package main 

import "fmt"

func main(){


	var side1,side2,hypotenuse,count,sidesSquared,hypotenuseSquared int;

	count = 0;

	
	//fmt.Printf("side1\tside2\thypotenuse")

	for side1 = 1; side1 <= 500; side1++ {

		for side2 = 1; side1 <= 500; side2++ {
			

			for hypotenuse = 1; hypotenuse <= 500; hypotenuse++ {

				sidesSquared = (side1 * side1) + (side2 * side2)

	            hypotenuseSquared = hypotenuse * hypotenuse


               if( hypotenuseSquared == sidesSquared) {


                 fmt.Printf("%d\t%d\t%d",side1,side2,hypotenuse)

                 count++

               }



			}

		}
	}
}