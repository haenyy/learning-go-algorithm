package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 시저 암호
* @Author haenyy
* @Method solution
* @Param s string, n int
* @Return answer int
* @Other 문자열 s와 거리 n을 입력받아 s를 n만큼 민 암호문을 만드는 함수 solution 완성하기
 */

func main() {
	fmt.Println("answer?", solution("AB", 1))
}

func solution(s string, n int) string {
	str := []rune(s)
	offset := rune(n)

	for i, value := range str {

		if 'A' <= value && value <= 'Z' { //대문자
			str[i] = ((value + offset - 'A') % 26) + 'A'
			fmt.Println(((value + offset - 'A') % 26) + 'A')
		} else if 'a' <= value && value <= 'z' { //소문자
			str[i] = ((value + offset - 'a') % 26) + 'a'
		} else { //공백일 때
			str[i] = value
		}
	}
	return string(str)
}
