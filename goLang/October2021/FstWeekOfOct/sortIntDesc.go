package main

import (
	"fmt"
	"sort"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 정수 내림차순으로 배치하기
* @Author haenyy
* @Method solution
* @Param n int64
* @Return int64
* @Other n의 각 자릿수를 큰 것부터 작은 순으로 정렬한 새로운 정수를 리턴 ex > 118372 -> 873211 리턴
 */

func main() {
	fmt.Println("answer?", solution(1234))
}

func solution(n int64) int64 {
	var num []int
	var answer int64 = 0

	for n > 0 { //숫자의 자릿수별로 나눠서 num 배열에 더하기
		num = append(num, int(n%10))
		n /= 10
	}

	sort.Slice(num, func(u, v int) bool {
		return num[u] > num[v]
	})

	for _, value := range num {
		answer = answer*10 + int64(value)
	}

	return answer
}
