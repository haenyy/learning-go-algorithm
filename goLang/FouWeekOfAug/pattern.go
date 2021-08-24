package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 수박수박수박수?
* @Author haenyy
* @Method solution
* @Param n int
* @Return string
* @Other 길이가 n이고 "수박수박수..."와 같은 패턴을 유지하는 문자열을 리턴
 */

func main() {

	fmt.Println("answer?", solution(3))
}

func solution(n int) string {

	answer := ""

	for i := 0; i < n; i++ {

		if i%2 == 0 {
			answer += "수"

		} else {
			answer += "박"

		}
	}

	return answer
}
