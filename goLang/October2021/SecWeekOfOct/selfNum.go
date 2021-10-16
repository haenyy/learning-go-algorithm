package main

import (
	"fmt"
)

/**
* @Title ◆ BACKJOON > 단계별 > 셀프넘버
* @Author haenyy
* @Method solution
* @Param n int
* @Return checkSelfNum map[int]bool
* @Other 생성자가 없는 숫자를 셀프 넘버라고 한다. 100보다 작은 셀프 넘버는 총 13개가 있다.
* 		 1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, 97
* 		 10000보다 작거나 같은 셀프 넘버를 한 줄에 하나씩 출력하는 프로그램을 작성.
 */

func main() {
	selfNum := solution(10000)
	for i := 1; i < len(selfNum); i++ {
		if selfNum[i] == false {
			fmt.Println(i)
		}
	}
}

func solution(n int) (checkSelfNum map[int]bool) {

	checkSelfNum = make(map[int]bool, n+1) //셀프 넘버 초기화

	for i := 0; i < n+1; i++ {
		checkSelfNum[i] = false
	}

	for i := 0; i < n+1; i++ {
		var sum = i
		var number = i
		for j := number; j != 0; j /= 10 {
			sum += j % 10
		}

		if sum <= n {
			checkSelfNum[sum] = true
		}
	}
	return checkSelfNum
}
