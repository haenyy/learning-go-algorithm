package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
* @Title ◆ BACKJOON > 단계별 > 1193.분수찾기
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other x번째 분수를 구하는 프로그램 작성
 */

func main() {
	var x int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &x)

	var level = 1
	var currentNum = 1

	for true {
		if currentNum > x {
			break
		}
		level++

		for i := level; i > 0; i-- {
			currentNum++

			if currentNum == x {
				if level%2 == 1 {
					fmt.Println(strconv.Itoa(i) + "/" + strconv.Itoa(level+1-i))
				} else {
					fmt.Println(strconv.Itoa(level+1-i) + "/" + strconv.Itoa(i))
				}
			}
		}
	}
}
