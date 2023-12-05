package main
import "fmt"

var g int = 10 //package global variable

func main(){
	var m int = 20 //func local variable

	{
		var s int = 50 //bracket local variable
		fmt.Println(m, s, g)
	} //s variable removed

	m = s + 20 //Error ocurred, undefined s
}