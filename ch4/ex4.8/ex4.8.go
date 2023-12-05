package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b //real value is 4266663.334329, but present 4266663 because of size limit
	var d float32 = c * 3 //real value is 12799990.002987, but present 127999890 because of size limit

	fmt.Println(a, b, c, d)
}