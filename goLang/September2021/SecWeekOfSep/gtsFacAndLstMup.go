package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 최대공약수와 최소공배수
* @Author haenyy
* @Method solution
* @Param n int, m int
* @Return []int
* @Other 두 수를 입력받아 두 수의 최대공약수와 최소공배수를 반환
 */

func main() {
	fmt.Println("answer?", solution(2, 5))
}

func solution(n int, m int) []int {

	answer := []int{}

	//gcd: 최대공약수, gcm: 최소공배수
	gcd := gcdFunc(n, m)
	gcm := n * m / gcd

	answer = []int{gcd, gcm}

	return answer
}

//최대공약수 return
func gcdFunc(a int, b int) int {

	for a > 0 {
		a, b = b%a, a
		fmt.Println(a, b)
	}

	return b
}
