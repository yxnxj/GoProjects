package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	stdin := bufio.NewReader(os.Stdin) // 입력 스트림에서 한 줄을 읽는 Reader 객체 stdin에 할당

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 에러 발생 후 스트림에서 \n이 나타날 때 까지 읽어서 스트림 비우기
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b) // 입력 다시 받기

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}