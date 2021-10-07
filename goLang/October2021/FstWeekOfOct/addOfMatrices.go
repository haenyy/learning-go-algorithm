package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 행렬의 덧셈
* @Author haenyy
* @Method solution
* @Param arr1 [][]int, arr2 [][]int
* @Return [][]int
* @Other 2개의 행렬 arr1과 arr2를 입력받아, 행렬 덧셈의 결과를 반환
 */

func main() {
	var arr1 = [][]int{{1, 2}, {2, 3}}
	var arr2 = [][]int{{1}, {2}}

	fmt.Println("answer?", solution(arr1, arr2))

}

func solution(arr1 [][]int, arr2 [][]int) [][]int {

	for z := range arr1 {
		for x := range arr1[0] {
			arr1[z][x] += arr2[z][x]
		}
	}

	return arr1
}
