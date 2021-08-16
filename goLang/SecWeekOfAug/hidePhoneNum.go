package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 핸드폰 번호 가리기
* @Author haenyy
* @Method solution
* @Param phone_number string
* @Return answer string
* @Other 전화번호의 뒷 4자리를 제외한 나머지 숫자를 전부 '*' 로 가린 문자열을 리턴
 */

func main() {
	fmt.Println("answer?", solution("01012345678"))
}

func solution(phone_number string) string {
	answer := ""
	pNum := []rune(phone_number) //rune type > 유니코드 제어 타입 (유니코드(UTF-8)을 표현하는 타입)

	loopNum := len(phone_number)
	for i := 0; i < loopNum-4; i++ {
		pNum[i] = '*'
	}

	answer = string(pNum)

	return answer
}
