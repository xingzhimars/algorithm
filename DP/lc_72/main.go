package main

// dp[i][j]：以w1[i]，w2[j]为结尾，w1转换成w2所使用的最少操作数
// w1[i] == w2[j]，不处理
// w1[i] != w2[j]，
//      w1[i-1] == w2[j-1]，插入w1[i]、w2[j]字符，操作数+1，dp[i][j] = dp[i-1][j-1] + 1
//      w1[i-1] == w2[j]，删除w1[i]字符，操作数+1，dp[i][j] = dp[i-1][j] + 1
//      w1[i] == w2[j-1]，插入w1[i+1]字符与w2[j]一样，操作数+1，dp[i][j] = dp[i][j-1] + 1

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 将w1字符变成w2，w2字符为空，只要每次删除一个字符即可，有几个字符，操作数就是几
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}

	// 同上
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}

	// dp[i][j]对应的是w1[i-1]w2[j-1]
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
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
