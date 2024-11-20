package main

// dp[i]：表示整数i，拆分成k个正整数的和后的乘积
// 有一个整数j，小于i，可以拆分成两种情况：
// 1. 拆分成两个整数，j * (i-j)
// 2. 拆分成大于2个的整数，即将 i-j的结果继续拆分，dp[i-j] * j
func integerBreak(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
