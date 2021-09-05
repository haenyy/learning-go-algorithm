package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 자연수 뒤집어 배열로 만들기
* @Author haenyy
* @Method solution
* @Param n int64
* @Return []int
* @Other 자연수 n을 뒤집어 각 자리 숫자를 원소로 가지는 배열 형태로 리턴
 */

func main() {
	fmt.Println("answer?", solution(12345))
}

func solution(n int64) []int {

	answer := []int{}
	var num []int

	for n > 0 {
		num = append(num, int(n%10))
		n /= 10
	}

	answer = num

	return answer
}
