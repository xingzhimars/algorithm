package main

const N = 110

var dp [N][N]int

func uniquePaths(m int, n int) int {
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 路径，不是步数，因此不需要+1
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
