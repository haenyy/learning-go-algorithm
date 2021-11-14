package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 1712.손익분기점
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other a, b, c가 주어졌을 때, 손익분기점 출력, 손익분기점이 존재하지 않으면 -1을 출력
 */

func main() {
	var a, b, c float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a, &b, &c)

	if c-b <= 0 {
		fmt.Println(-1)
		return
	}

	roundUp := int(math.Ceil(a / (c - b))) //Ceil > 실수-정수 반올림

	if roundUp == int(a/(c-b)) {
		fmt.Println(roundUp + 1)
	} else {
		fmt.Println(roundUp)
	}
}
