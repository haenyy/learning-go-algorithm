package main

import "fmt"

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 가운데 글자 가져오기
* @Author haenyy
* @Method solution
* @Param s string / 문자열
* @Return answer
* @Other 단어의 길이가 짝수일 때 가운데 두글자 반환, 홀수일 때 가운데 한글자 반환
 */
func solution(s string) string {

	answer := ""
	even := len(s)/2 - 1
	odd := len(s) / 2

	if len(s)%2 == 0 {
		answer = string(s[even : even+2]) //sub slice 적용
	} else {
		answer = string(s[odd])
	}

	fmt.Println("even?", even)
	fmt.Println("odd?", odd)
	fmt.Println("answer?", answer)

	return answer
}

func main() {
	solution("qwer") //test case : "abcde", "qwer"
}
