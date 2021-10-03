package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 약수의 합
* @Author haenyy
* @Method solution
* @Param n int
* @Return int
* @Other 정수 n을 입력받아 n의 약수를 모두 더한 값을 리턴
 */

func main() {
	n := 12
	fmt.Println("answer?", solution(n))

}

func solution(n int) int {

	answer := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			answer += i
		}
	}
	return answer
}
