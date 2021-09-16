package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 짝수와 홀수
* @Author haenyy
* @Method solution
* @Param num int
* @Return string
* @Other 정수 num이 짝수이면 Even 반환, 홀수일 때 Odd 반환
 */

func main() {
	fmt.Println("answer?", solution(3))
}

func solution(num int) string {

	answer := ""

	if num%2 == 0 {
		answer = "Even"
	} else {
		answer = "Odd"
	}

	return answer
}
