package main

// dp[j]：凑成金额为j的组合数
// 假设当前有一个coins[i]硬币，可用可不用，dp[j] = dp[j] + dp[j-coins[i]]，求的是方案数
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)

	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}

	return dp[amount]
}
