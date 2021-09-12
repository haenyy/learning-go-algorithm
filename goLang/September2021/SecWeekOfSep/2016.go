package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 2016년
* @Author haenyy
* @Method solution
* @Param a int, b int
* @Return stirng
* @Other 2016년 a월 b일은 무슨 요일인지 리턴
 */

func main() {
	a := 5
	b := 24
	fmt.Println("answer?", solution(a, b))
}

func solution(a int, b int) string {

	day := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	days := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	answer := ""
	sum := 0

	for i := 0; i < a-1; i++ {
		sum += days[i]
		fmt.Println(sum)
	}

	sum += b + 4 //달에 포함되는 일수만큼 더하고 1월 1일이 금요일이므로 4를 추가로 더해줌
	fmt.Println(sum)

	answer = day[sum%7]

	return answer
}
