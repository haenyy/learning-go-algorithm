package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 평균 구하기
* @Author haenyy
* @Method solution
* @Param arr []int
* @Return float64
* @Other 정수를 담고 있는 배열 arr의 평균값을 return
 */

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("answer?", solution(arr))
}

func solution(arr []int) float64 {

	answer := 0.0

	var length int
	length = len(arr)
	// for i := 0; i <= length; i++ {
	// 	answer += arr[i]
	// }

	sum := 0
	for _, value := range arr {
		sum += value
	}

	answer = float64(sum) / float64(length)

	return answer
}
