package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 11654.ASCII Code
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 알파벳 소문자, 대문자, 숫자 0-9 중 하나가 주어졌을 때 아스키 코드값을 출력
 */

func main() {
	var input byte

	reader := bufio.NewReader(os.Stdin) //bufio 입출력 인터페이스

	input, _ = reader.ReadByte() //문자열 읽기

	fmt.Println("%d\n", input)
}
