package main

import (
	"fmt"
	"unicode"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 이상한 문자 만들기
* @Author haenyy
* @Method solution
* @Param s string
* @Return string
* @Other 문자열 s에서 짝수번째 알파벳은 대문자로 홀수번째 알파벳은 소문자로 바꾼 문자열 return
 */

func main() {
	fmt.Println("answer?", solution("try hello world"))
}

func solution(s string) string {

	answer := ""

	strs := []rune(s)
	idx := 0

	for i, value := range strs {
		//IsSpace로 배열 내 공백 문자열을 찾아서 공백이면 true
		if unicode.IsSpace(value) {
			idx = i + 1
		} else if (i-idx)%2 == 0 { //짝수일 때 대문자로
			strs[i] = unicode.ToUpper(strs[i])
		} else { //홀수일 때 소문자로
			strs[i] = unicode.ToLower(strs[i])
		}
	}
	answer = string(strs)

	return answer
}
