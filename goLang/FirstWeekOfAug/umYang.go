package main

import (
	"fmt" //import fmt, main() 을 포함한 패키지
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 월간 코드 챌린지 시즌2 > 음양 더하기
* @Author haenyy
* @Method solution
* @Param absolutes []int, signs []bool / 정수의 절대값 배열, 정수들의 부호를 담은 배열
* @Return answer / 실제 정수들의 합
* @Other 실제 정수들의 합을 구하여 return 하는 solution 함수 만들기
 */
func solution(absolutes []int, signs []bool) int {

	//return 초기화
	answer := 0

	for i := 0; i < len(absolutes); i++ {
		if signs[i] == true { //양수일 때
			answer += absolutes[i]
		} else { //음수일 때
			answer -= absolutes[i]
		}
	}
	return answer
}

func main() {
	var a [3]int = [3]int{4, 7, 12}
	var b = [3]bool{true, false, true}
	var c [3]int = [3]int{1, 2, 3}
	var d = [3]bool{false, false, true}

	fmt.Println("arraysize", len(c), len(d))
	fmt.Println("return", solution(a[:], b[:]))

}
