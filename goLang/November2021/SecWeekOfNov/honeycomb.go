package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 2292.벌집
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 입력으로 주어진 방까지 최소 개수의 방을 지나서 갈 때 몇 개의 방을 지나는지 출력
 */

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var currentNum = 1
	var room = 1

	for true {
		if currentNum >= n {
			break
		}

		currentNum += 6 * room
		room++
	}

	fmt.Println(room)
}
