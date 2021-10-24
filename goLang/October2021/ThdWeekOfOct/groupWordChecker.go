package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
* @Title ◆ BACKJOON > 단계별 > 1316.그룹단어체커
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 단어 N개를 입력받아 그룹 단어의 개수를 출력 (그룹단어 : 단어에 존재하는 모든 문자에 대해서, 각 문자가 연속해서 나타나는 경우)
 */

func main() {
	var n int

	reader := bufio.NewReader(os.Stdin) //bufio 입출력 인터페이스

	fmt.Fscanln(reader, &n)

	var word string
	var groupCount int

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &word)

		var letters = make(map[uint8]bool)
		var prev uint8
		var isGroup = true

		for j := 0; j < len(word); j++ {
			_, exist := letters[word[j]]

			if !exist {
				letters[word[j]] = true

			} else if prev != word[j] {
				isGroup = false
				break

			}
			prev = word[j]
		}

		if isGroup {
			groupCount++
		}

	}

	fmt.Println(groupCount)
}
