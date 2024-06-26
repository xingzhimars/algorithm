package main

// 思路：dp
// dp[i][j]：考虑word1前i个字符，word2前j个字符，最少操作数
// 当word1[i] == word2[j]，不需要操作
// 当word1[i] != word2[j]
// 1. word1[i-1] == word2[j-1]，替换    dp[i-1][j-1] + 1
// 2. word1[i-1] == word2[j]，删除word1[i] dp[i-1][j] + 1
// 3. word1[i] == word2[j-1]，插入一个字符 dp[i][j-1] + 1

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j-1], dp[i-1][j])) + 1
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
