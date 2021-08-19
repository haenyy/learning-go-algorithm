package main

import (
	"fmt"
	"math"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 위클리 챌린지 > 1주차
* @Author haenyy
* @Method solution
* @Param price int, money int, count int
* @Return answer
* @Other 놀이기구를 count번 타게 되면 price가 얼마나 모자라는지 return
 */

func main() {
	fmt.Println("answer?", solution(3, 20, 4))
}

func solution(price int, money int, count int) int64 {

	answer := -1

	fee := 0
	for i := 1; i <= count; i++ {
		fee += price * i
	}

	//정확성 위해 절대값 처리 해줌. 근데 이게 정확한 처리가 맞을까? 고민해볼것.
	tFee := int(math.Abs(float64(fee)))
	tMoney := int(math.Abs(float64(money)))

	if fee > money { //금액 * 횟수가 처음 가지고 있던 금액보다 적다면
		answer = tFee - tMoney
	} else {
		answer = 0
	}

	return int64(answer)
}
