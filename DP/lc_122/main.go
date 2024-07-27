package main

// 可以多次买卖
// dp[i][0]: 第i天结束时，手中没有股票的情况下，最大利润
// dp[i][1]: 第i天结束时，手中有股票的情况下，最大利润
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], prices[i]+dp[i-1][1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}
