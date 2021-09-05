package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 직사각형 별찍기
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other n, m을 입력받아 가로 길이가 n, 세로 길이가 m인 직사각형 형태 출력
 */

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
