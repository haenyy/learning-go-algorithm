package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 콜라츠 추측
* @Author haenyy
* @Method solution
* @Param num int
* @Return answer int
* @Other 짝수/홀수에 따라 곱셈 및 나눗셈을 달리하여 결과가 1이 될때까지 반복. 단 500번 넘을 경우 -1 리턴
 */

func main() {
	fmt.Println("answer?", solution(626331))
}

func solution(num int) int {
	answer := 0

	cnt := 0

	for num > 1 && cnt < 500 {
		if num%2 == 0 { //짝수라면 2로 나누고
			num = num / 2
		} else { //홀수라면 3을 곱하기 1을 더함
			num = (num * 3) + 1
		}
		cnt++
	}
	if cnt == 500 { //500번 반복해도 1이 되지 않을 때 -1 리턴
		answer = -1
	} else {
		answer = cnt
	}

	return answer
}
