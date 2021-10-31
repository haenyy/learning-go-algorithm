package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 3052.나머지
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 각 숫자가 몇 번 나왔는지 저장하기 위해 일차원 배열을 만드는 문제
*		 10개를 입력받아 42로 나눈 나머지를 구하고 서로 다른 값이 몇 개 있는지 출력
 */

func main() {

	reader := bufio.NewReader(os.Stdin) //bufio 입출력 인터페이스
	var numbers = make([]int, 10)
	var mods []int

	for i := range numbers {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
		var mod = numbers[i] % 42

		if !contains(mods, mod) {
			mods = append(mods, mod)
		}

	}

	fmt.Println(len(mods))
}

func contains(numbers []int, number int) bool {
	//numbers = 입력한 숫자, number = 42로 나눈 나머지
	for _, n := range numbers {
		if number == n {
			return true
		}
	}
	return false
}
