package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b) // set minimum print size to 5 /  123,  456
	fmt.Printf("%05d, %05d\n", a, b) //if variable print size less than setting, fill with 0 / 00123,00456
	fmt.Printf("%-5d, %-05d\n", a, b) //sorting on left side / 123  ,456  

	// minimum size setting ignored, print size is bigger than setting
	fmt.Printf("%5d, %5d\n", c, c) 
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)
	
}