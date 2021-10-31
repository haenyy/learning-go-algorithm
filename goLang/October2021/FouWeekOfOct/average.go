package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 1546.평균
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 평균을 조작하는 문제 / 점수 중 최댓값(Max)을 골라 총점/Max*100 이 되도록 출력
 */

func main() {
	var n int
	var max, sum int

	reader := bufio.NewReader(os.Stdin) //bufio 입출력 인터페이스

	fmt.Fscanln(reader, &n)

	var scores = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &scores[i])
		sum += scores[i]

		if max < scores[i] {
			max = scores[i]
		}
	}

	fmt.Println(float64(sum) / float64(n) / float64(max) * 100.0)
}
