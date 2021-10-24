package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
* @Title ◆ BACKJOON > 단계별 > 2941.크로아티아알파벳
* @Author haenyy
* @Method main
* @Param
* @Return
* @Other 입력으로 주어진 단어가 몇 개의 크로아티아 알파벳으로 이루어져 있는지 출력
 */

func main() {
	var word string

	reader := bufio.NewReader(os.Stdin) //bufio 입출력 인터페이스

	fmt.Fscanln(reader, &word)

	word = strings.Replace(word, "c=", "!", -1) //n이 -1일 때 교차 제한 없음
	word = strings.Replace(word, "c-", "@", -1)
	word = strings.Replace(word, "dz=", "#", -1)
	word = strings.Replace(word, "d-", "$", -1)
	word = strings.Replace(word, "lj", "%", -1)
	word = strings.Replace(word, "nj", "^", -1)
	word = strings.Replace(word, "s=", "&", -1)
	word = strings.Replace(word, "z=", "*", -1)

	var alphbets []string

	for i := 0; i < len(word); i++ {
		alphbets = append(alphbets, string(word[i]))
	}

	fmt.Println(len(alphbets))
}
