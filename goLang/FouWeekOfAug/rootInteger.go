package main

import (
	"fmt"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 수박수박수박수?
* @Author haenyy
* @Method solution
* @Param n int64
* @Return int64
* @Other n이 어떤 양의 정수 x의 제곱인지 아닌지 판단 하여 맞다면 x+1의 제곱을 리턴, 아니라면 -1을 리턴
 */

func main() {

	fmt.Println("answer?", solution(121))
}

func solution(n int64) int64 {

	answer := 0

	//항상 int와 int64의 형변환이 어렵다 쉬워보이는데 코드 짤 때마다 헷갈리는...
	//처음에는 i에만 형변환을 해주었는데 곱셈 부분에서 잘 읽히지 않았던..?
	//그래서 아예 계산식에 int64 형변환을 적용했다.
	for i := 0; int64(i*i) <= n; i++ {
		if int64(i*i) == n {
			answer = (i + 1) * (i + 1)
		} else {
			answer = -1
		}
	}

	return int64(answer)
}
