package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 최댓값과 최솟값
* @Author haenyy
* @Method solution
* @Param s string
* @Return string
* @Other 문자열 s에 나타나는 숫자 중 최솟값과 최댓값을 찾아 반환
 */

func main() {
	fmt.Println("answer?", solution("-1 -2 -3 -4"))
}

func solution(s string) string {

	answer := ""
	slice := strings.Split(s, " ")
	fmt.Println(slice)

	min := slice[0]
	max := slice[0]
	var strToN int
	var minToN int
	var maxToN int
	for _, str := range slice {

		if n, err := strconv.Atoi(str); err == nil {
			strToN = n
		}

		if n, err := strconv.Atoi(min); err == nil {
			minToN = n
		}
		if n, err := strconv.Atoi(max); err == nil {
			maxToN = n
		}

		if strToN < minToN {
			minToN = strToN
		}

		if strToN > maxToN {
			maxToN = strToN
		}
	}
	answer = strconv.Itoa(minToN) + " " + strconv.Itoa(maxToN)

	return answer
}
