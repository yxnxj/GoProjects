package main
import "fmt"

func main(){
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	// fmt.Scan(&a, &b)
	// fmt.Println(a, b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}