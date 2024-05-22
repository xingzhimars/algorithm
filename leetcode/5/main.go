package main

// dp
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				dp[i][j] = true
			}
		}
	}

	start, end := 0, 0
	maxLen := 0
	for k := 1; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}
			if dp[i][j] {
				if k+1 > maxLen {
					maxLen = k + 1
					start = i
					end = j
				}
			}
		}
	}
	return s[start : end+1]
}
