package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 자릿수 더하기
* @Author haenyy
* @Method solution
* @Param n int
* @Return int
* @Other 자연수 n이 주어지면 n의 각 자릿수의 합을 구해서 리턴
 */

func main() {
	n := 987
	fmt.Println("answer?", solution(n))

}

func solution(n int) int {

	answer := 0

	for n > 0 {
		answer += n % 10
		n /= 10
	}

	return answer
}
