package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	splits := strings.Split(s, " ")
	i, j := 0, len(splits)-1
	for i <= j {
		splits[i], splits[j] = splits[j], splits[i]
		i++
		j--
	}
	res := ""
	for i := 0; i < len(splits); i++ {
		if splits[i] == "" {
			continue
		}
		res += splits[i]
		res += " "
	}
	res = strings.TrimRight(res, " ")
	return res
}

func reverse(s string) {

}

func main() {
	s := "a good   example"
	_ = reverseWords(s)
	//fmt.Println(words)
}
