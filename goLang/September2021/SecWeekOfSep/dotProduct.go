package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 월간 코드 챌린지 시즌1 > 내적
* @Author haenyy
* @Method solution
* @Param a []int, b []int
* @Return int
* @Other a와 b의 내적을 return (a[0] * b[0] + a[1] * b[1] ··· a[n-1] * b[n-1])
 */

func main() {
	arr1 := []int{-3, -1, 0, 2}
	arr2 := []int{1, 2, 3, 4}
	fmt.Println("answer?", solution(arr1, arr2))
}

func solution(a []int, b []int) int {

	answer := 0

	for i, value := range a {
		answer += value * b[i]
	}

	return answer
}
