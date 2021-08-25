package main

import (
	"fmt"
	"sort"
)

/**
* @Title ◆ Programmers > 코딩테스트 연습 > 연습문제 > 나누어 떨어지는 숫자 배열
* @Author haenyy
* @Method solution
* @Param arr []int, divisor int
* @Return []int
* @Other arr의 요소 중 divisor로 나누어 떨어지는 값을 오름차순으로 정렬한 배열을 반환
 */

func main() {
	arr := []int{5, 9, 7, 10}
	fmt.Println("answer?", solution(arr, 5))
}

func solution(arr []int, divisor int) []int {
	var answer []int
	fmt.Println(len(arr))
	// for i := 0; i < len(arr); i++ {

	// 	fmt.Println("arr / divisor ?", arr[i]%divisor)

	// 	if arr[i]%divisor == 0 {
	// 		//answer[i] += arr[i]
	// 	} else {
	// 		answer[i] = -1
	// 		return answer
	// 	}
	// }

	for _, value := range arr {
		if value%divisor == 0 {
			answer = append(answer, value)
		}
	}

	//오름차순 정렬을 위한 익명함수 사용. 슬라이스의 정렬 (Slice Sort)
	//참고 : http://www.gisdeveloper.co.kr/?p=5712

	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})

	//for문 안에서 분기태워서 넣으려했더니 배열 index가 범위밖을 벗어나는 오류가 났다
	//애초에 조건에 해당하는 것만 for문 돌리고 거르는 작업을 for문 밖에서 해줘도 되겠다.
	if len(answer) == 0 {
		return []int{-1}
	}

	return answer
}
