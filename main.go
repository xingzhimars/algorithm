package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		s := strs[i]
		prefix = getPrefix(prefix, s)
	}

	return prefix
}

func getPrefix(s1, s2 string) string {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			return s1[:i]
		}
		i++
		j++
	}
	return ""
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
