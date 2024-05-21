package main

import "fmt"

// kmp算法
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	i, j := 0, 0
	for i < n && j < m {
		if haystack[i] == needle[j] {
			j++
		} else {
			j = 0
		}
		i++
		if j == m {
			return i - j
		}
	}
	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issip"
	res := strStr(haystack, needle)
	fmt.Println(res)
}
