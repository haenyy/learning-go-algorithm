package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 2884.알람 시계
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 정수 n개가 주어졌을 때, n개의 합을 구하는 함수 출력
 */

func main() {
	var h, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &h, &m)

	if m-45 >= 0 {
		fmt.Printf("%d %d\n", h, m-45) //45분보다 크면 시, 분 그대로
	} else if h == 0 { //12시라면 11시 + 분 계산
		fmt.Printf("%d %d\n", 23, 15+m)
	} else {
		fmt.Printf("%d %d\n", h-1, 15+m)
	}
}
