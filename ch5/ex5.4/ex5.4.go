package main
import "fmt"

func main(){
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"
	// /t is tab
	// to use " in string write like \"
	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}