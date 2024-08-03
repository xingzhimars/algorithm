package main

func rob(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = nums[0]

	// 1. 必选i，那么i-1必定不能选，则dp[i][1] = dp[i-1][0] + nums[i]
	// 2. 必不选i，那么i-1可选可不选，则dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[n-1][0], dp[n-1][1])
}
