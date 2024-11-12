package main

// dp[i][0]：第i天结束时，手上有股票
// dp[i][1]：第i天结束时，手上没股票，不处于冷冻期
// dp[i][2]：第i天结束时，手上没股票，处于冷冻期

// 状态转移方程：
// 1. 当天买入的；2. 昨天就有了
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[i])

// 1. 昨天不持股，处于冷冻期；2. 昨天不持股，不处于冷冻期
// dp[i][1] = max(dp[i-1][2], dp[i-1][1])

// 今天卖的
// dp[i][2] = dp[i-1][0]+prices[i]
func maxProfit(prices []int) int {
	n := len(prices)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][2], dp[i-1][1])
		dp[i][2] = dp[i-1][0] + prices[i]
	}

	return max(dp[n-1][1], dp[n-1][2])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
