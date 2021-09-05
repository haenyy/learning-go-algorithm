package main

import (
	"fmt"
	"strconv"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 문자열을 정수로 바꾸기
* @Author haenyy
* @Method solution
* @Param s string
* @Return int
* @Other 문자열 s를 숫자로 변환한 결과를 반환
 */

func main() {
	fmt.Println("answer?", solution("-1234"))
}

func solution(s string) int {

	answer := 0

	num, _ := strconv.Atoi(s)

	answer = num
	return answer
}
