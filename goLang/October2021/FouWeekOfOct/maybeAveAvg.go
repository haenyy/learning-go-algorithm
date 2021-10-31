package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 4344.평균은 넘겠지
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 단어 N개를 입력받아 그룹 단어의 개수를 출력 (그룹단어 : 단어에 존재하는 모든 문자에 대해서, 각 문자가 연속해서 나타나는 경우)
 */

func main() {
	var c int

	reader := bufio.NewReader(os.Stdin) //bufio 입출력 인터페이스
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &c) //테스트 케이스 개수

	defer writer.Flush() //defer - 지연처리 (참고: http://golang.site/go/article/20-Go-defer%EC%99%80-panic)

	for i := 0; i < c; i++ { //학생 수 입력받고
		var n int
		fmt.Fscanf(reader, "%d ", &n)

		var scores = make([]float64, n)
		var sum, avg float64

		for j := 0; j < n; j++ { //입력받은 학생 수만큼 점수 입력
			fmt.Fscanf(reader, "%f ", &scores[j])
			sum += scores[j]
		}

		avg = sum / float64(n)

		var proportion float64
		for _, val := range scores { //평균을 넘는 학생들의 비율 계산, 소숫점 셋째 자리까지 출력
			if val > avg {
				proportion += (1 / float64(n))
			}
		}

		fmt.Fprintf(writer, "%.3f%s\n", math.Round(proportion*100000)/1000, "%")
	}
}
