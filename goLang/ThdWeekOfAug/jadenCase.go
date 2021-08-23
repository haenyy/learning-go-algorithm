package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > JadenCase 문자열 만들기
* @Author haenyy
* @Method solution
* @Param s string
* @Return string
* @Other 문자열 s가 주어졌을 때 s를 JadenCase로 바꾼 문자열을 리턴 (※ JadenCase : 모든 단어의 첫 문자가 대문자, 그 외의 알파벳은 소문자)
 */

func main() {
	fmt.Println("answer?", solution("foR 1the bEst world"))
}

func solution(s string) string {

	answer := ""
	slice := strings.Split(s, " ")
	fmt.Println(slice)

	for _, str := range slice {

		lower := strings.ToLower(str)
		tmp := []rune(lower)

		tmp[0] = unicode.ToTitle(tmp[0])

		answer += string(tmp)
		answer += " "

	}
	answer = strings.TrimRight(answer, " ")

	return answer
}
