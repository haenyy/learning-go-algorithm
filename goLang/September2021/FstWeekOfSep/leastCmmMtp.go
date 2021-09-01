package main

import (
	"fmt"
	"math"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > N개의 최소공배수
* @Author haenyy
* @Method solution
* @Param arr []int
* @Return int
* @Other n개의 숫자를 담은 배열 arr이 입력되었을 때 이 수들의 최소공배수를 반환
 */

func main() {
	//arr := []int{2, 4, 6, 8, 14}
	arr := []int{1, 2, 3}
	fmt.Println("answer?", solution(arr))
}

func solution(arr []int) int {

	answer := arr[0]

	//유클리드 호제법 이용
	//최대공약수 : 두 수일 때 한 수(b)를 나눠서 나머지를 임시변수(tmp)에 저장하고
	//나누어진 수(a)는 다시 나눈 수(b)가 되고 나눈 수(b)는 임시변수(r)이 되어
	//0이 아닐때까지 반복하다 0이 되면 a를 리턴하게 되는데 a가 a, b의 최대공약수가 됨)
	//최소공배수 : (두 수의 곱 / 두 수의 최대공약수)
	//2개의 최소공배수를 구한 다음 그 값과 계산되지 않은 값들 중 1개를 선택해 다시 최소공배수 구하기

	var a, b, r int

	for i := 1; i < len(arr); i++ {
		//get gcd(최대공약수)
		a = int(math.Max(float64(answer), float64(arr[i])))
		b = int(math.Min(float64(answer), float64(arr[i])))

		for (a % b) != 0 {
			r = a % b
			a = b
			b = r
		}

		answer = answer * arr[i] / b
	}

	return answer
}
