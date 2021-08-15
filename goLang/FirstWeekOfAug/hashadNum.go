package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 하샤드 수
* @Author haenyy
* @Method solution
* @Param x int / 정수
* @Return answer
* @Other 양의 정수 x가 하샤드 수이려면 x의 자릿수의 합으로 x가 나누어져야 함
 */

func main() {
	a := 13
	b := solution(a)
	fmt.Println("x?", b)

}

func solution(x int) bool {

	sum := 0 //각 자릿수 합계
	num := x //각 자릿수의 합을 계산하기 위해 사용할 변수
	answer := false

	for num/10 != 0 { //한 자릿수가 나올때까지
		sum += num % 10 //10으로 나눈 몫을 합계에 더해준다
		num /= 10       //다음 자릿수 구하기 위해 나눠주기
	}
	sum += num //마지막에 남는 몫이 마지막 자릿수이기 때문에 더해준다

	if x%sum == 0 {
		answer = true
	}

	return answer
}
