package main

// 完全背包问题
func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)

	// dp的初始化，注意dp[0]为0
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}

		}
	}

	if dp[amount] < amount+1 {
		return dp[amount]
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
