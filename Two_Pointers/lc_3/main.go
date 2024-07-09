package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	i, j := 0, 0
	set := make(map[byte]struct{})

	maxLen := 0

	for j < n {
		_, ok := set[s[j]]
		if !ok {
			set[s[j]] = struct{}{}
			j++
			maxLen = max(maxLen, j-i)
		} else {
			delete(set, s[i])
			i++
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
