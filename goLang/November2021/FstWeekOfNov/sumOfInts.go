package main

import "fmt"

/**
* @Title ◆ BACKJOON > 단계별 > 15596.정수 N개의 합
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 정수 n개가 주어졌을 때, n개의 합을 구하는 함수 출력
 */

func main() {
	var arr = []int{5, 6, 7, 8}
	fmt.Println(sum(arr))
}

func sum(a []int) int {
	var r int

	for _, val := range a {
		r += val
	}

	return r
}
