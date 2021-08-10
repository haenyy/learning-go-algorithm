package main

import (
	"fmt"
	"strconv"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 서울에서 김서방 찾기
* @Author haenyy
* @Method solution
* @Param seoul []string
* @Return answer
* @Other String형 배열 seoul의 element중 "Kim"의 위치 x를 찾아
* 		 "김서방은 x에 있다"는 String을 return 하는 solution 함수 만들기
 */
func solution(seoul []string) string {
	num := 0

	for i := 0; i < len(seoul); i++ {
		if seoul[i] == "Kim" {
			num = i
		}
	}
	//strconv > 숫자를 문자열로 변환
	answer := "김서방은 " + strconv.Itoa(num) + "에 있다"

	return answer
}

func main() {
	var d []string = []string{"Jane", "Kim"}
	fmt.Println(solution(d))
}
