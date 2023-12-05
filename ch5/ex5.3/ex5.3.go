package main
import "fmt"

func main(){
	var a = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n", a) // minimum size 8, print two nums below point, fill 0
	fmt.Printf("%08.2g\n", a) // minimum size 8, print two nums, fiil 0
	fmt.Printf("%8.5g\n", a) //minimum size 8, print five nums
	fmt.Printf("%f\n", c) //print six nums below point 
}