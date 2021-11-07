package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 15596.정수 N개의 합
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 정수 n개가 주어졌을 때, n개의 합을 구하는 함수 출력
 */

func main() {
	var n, cycle int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var newN = n
	for true {
		cycle++

		var sum int
		if newN/10 == 0 {
			sum = newN
		} else {
			sum = newN/10 + newN%10
		}

		newN = newN%10*10 + sum%10

		if newN == n {
			break
		}
	}

	fmt.Println(cycle)
}
