package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 피보나치 수
* @Author haenyy
* @Method solution
* @Param n int
* @Return int
* @Other 2 이상의 n이 입력되었을 때, n번째 피보나치 수를 1234567으로 나눈 나머지를 리턴
 */

func main() {
	n := 3
	fmt.Println("answer?", solution(n))
}

func solution(n int) int {

	f := 0
	s1 := 0
	s2 := 1
	result := 0

	for i := 1; i < n; i++ { //i가 0부터 시작하는 것, 1부터 시작하는 것 result 값 다름... 왜....?
		var temp = s1 + s2
		f = temp
		s1 = s2
		s2 = temp

		if s1 >= 1234567 {
			s1 %= 1234567
		}
		if s2 >= 1234567 {
			s2 %= 1234567
		}
	}

	result = f % 1234567

	return result
}
