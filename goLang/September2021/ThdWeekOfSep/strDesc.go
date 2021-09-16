package main

import (
	"fmt"
	"sort"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 문자열 내림차순으로 배치하기
* @Author haenyy
* @Method solution
* @Param s string
* @Return string
* @Other 문자열 s에 나타나는 문자를 큰것부터 작은 순으로 정렬해 새로운 문자열 리턴
 */

func main() {
	fmt.Println("answer?", solution("Zbcdefg"))
}

func solution(s string) string {

	answer := ""
	sbyte := []byte(s)

	sort.Slice(sbyte, func(i, j int) bool {
		return sbyte[i] > sbyte[j]
	})

	answer = string(sbyte)

	return answer
}
