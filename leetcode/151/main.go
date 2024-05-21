package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	n := len(ss)
	i, j := 0, n-1
	for i <= j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	var res string
	for i := 0; i < n; i++ {
		if ss[i] == "" {
			continue
		}
		res += ss[i]
		res += " "
	}
	res = strings.TrimRight(res, " ")
	return res
}

func main() {
	//s := "the sky is blue"
	//s := "a good   example"
	s := "  hello world  "
	words := reverseWords(s)
	fmt.Println(words)
}
