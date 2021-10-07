package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 제일 작은 수 제거하기
* @Author haenyy
* @Method solution
* @Param arr []int
* @Return []int
* @Other 정수를 저장한 배열, arr에서 가장 작은 수를 제거한 배열을 리턴. 단, 리턴하려는 배열이 빈 배열일 경우엔 배열에 -1 채워 리턴
 */

func main() {
	var arr = []int{4, 3, 2, 1}

	fmt.Println("answer?", solution(arr))

}

func solution(arr []int) []int {

	answer := []int{}

	if len(arr) == 1 {
		return []int{-1}
	} else {
		answer = removeIndex(arr, findMinIndex(arr))
	}
	return answer
}

/**
*	@Title 찾은 최솟값의 인덱스를 제거
**/
func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

/**
*	@Title 최솟값의 인덱스 찾기
**/
func findMinIndex(arr []int) int {
	index := 0
	for i, value := range arr {
		if arr[index] > value {
			index = i
		}
	}
	return index
}
