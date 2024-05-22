package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = getPrefix(prefix, strs[i])
	}
	return prefix
}

func getPrefix(s1, s2 string) string {
	i, j := 0, 0
	n, m := len(s1), len(s2)

	for i < n && j < m {
		if s1[i] != s2[j] {
			break
		}
		i++
		j++
	}
	return s1[:i]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
