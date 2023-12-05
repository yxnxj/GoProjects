package main
import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a) // int16 to int8 conversion
						 // overflow occured

	fmt.Println(a)
	fmt.Println(c)


}