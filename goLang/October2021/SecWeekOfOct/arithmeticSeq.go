package main

import (
	"fmt"
)

/**
* @Title ◆ BACKJOON > 단계별 > 1065.한수
* @Author haenyy
* @Method solution
* @Param a int, b int
* @Return int
* @Other N이 주어졌을 때, 1보다 크거나 같고, N보다 작거나 같은 한수(연속된 두 개의 수의 차이가 일정한 수열)의 개수를 출력하는 프로그램을 작성
 */

func main() {
	n := 110
	fmt.Println("answer?", solution(n))
}

func solution(num int) (count int) {

	if num < 100 {
		count = num
		return
	}

	for i := 100; i <= num; i++ {
		one := i % 10
		ten := i / 10 % 10
		hund := i / 100

		if hund-ten == ten-one {
			count++
		}
	}
	count += 99

	return count
}
