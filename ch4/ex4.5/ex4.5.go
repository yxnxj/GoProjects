package main

import "fmt"

func main(){
	a := 3 //int
	var b float64 = 3.5 //float
 
	var c int = int(b) //float to int conversion
	d := float64(a*c) //int to float conversion with product cal

	var e int64 = 7
	f := int64(d) * e //float to int conversion

	var g int = int(b * 3) //float to int conversion
	var h = int(b) * 3 //float to int conversion, different value with g
	
	fmt.Println(g, h, f)

}