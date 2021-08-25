package main

import (
	"fmt"
	"strconv"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 문자열 다루기 기본
* @Author haenyy
* @Method solution
* @Param s string
* @Return bool
* @Other 문자열 s가 숫자로만 구성되어있다면 True return 아니면 False return
 */

func main() {

	fmt.Println("answer?", solution("1234"))
}

func solution(s string) bool {

	answer := true

	//nil은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value를 의미함
	//명시적인 초기값 없이 선언된 변수는 zero value를 갖게 됨
	//(비어있는 것과 애초에 아무것도 없는 것은 다름)
	_, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		answer = false
	} else {
		answer = true
	}

	if len(s) != 4 && len(s) != 6 {
		answer = false
	}

	return answer
}
