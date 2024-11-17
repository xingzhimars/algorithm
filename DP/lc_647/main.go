package main

func countSubstrings(s string) int {
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

	for k := 1; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}
