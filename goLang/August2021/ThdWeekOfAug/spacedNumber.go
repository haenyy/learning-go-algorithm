package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > x만큼 간격이 있는 n개의 숫자
* @Author haenyy
* @Method solution
* @Param x int, n int
* @Return answer []int64
* @Other x부터 시작해 x씩 증가하는 숫자를 n개 지니는 리스트를 만드는 함수 solution 완성하기
 */

func main() {
	fmt.Println("answer?", solution(2, 5))
}

func solution(x int, n int) []int64 {

	var answer []int64
	var sum int64
	for i := 0; i < n; i++ {

		sum += int64(x)
		answer = append(answer, sum)

		// answer[i] = int64(x * (i + 1))
		// 이렇게 했을 때 에러가 나는데
		// go 언어의 배열은 인스턴스 값으로 초기화 할 수 없다고 한다
		// 즉, 고정 값을 가지는 상수로만 초기화가 가능하며
		// 크기를 자유롭게 변형 가능한 슬라이스를 이용해야 한다.
	}

	return answer
}
