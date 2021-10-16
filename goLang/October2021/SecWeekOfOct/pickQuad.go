package main

import (
	"fmt"
)

/**
* @Title ◆ BACKJOON > 단계별 > 사분면 고르기
* @Author haenyy
* @Method solution
* @Param a int, b int
* @Return int
* @Other 주어진 점이 어느 사분면에 속하는지 알아내어 리턴 (ex > (-12,5) 이면 2사분면 / (12,5) 이면 1사분면)
 */

func main() {
	a := 9
	b := -13
	fmt.Println("answer?", solution(a, b))
}

func solution(a int, b int) int {

	answer := 0

	if a > 0 {
		if b > 0 {
			answer = 1
		} else {
			answer = 4
		}

	} else {
		if b > 0 {
			answer = 2
		} else {
			answer = 3
		}
	}

	return answer
}
