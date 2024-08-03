package main

// 枚举两种情况
// 情况1：不选1，走一遍198流程
// 情况2：选1，走一遍198流程
// 两种情况取最大值

// dp[i][0]，从 1 ～ i中选，i不偷
// dp[i][1]，从1 ～ i中选，i偷
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// 不选1
	dp[0][0] = 0
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	res := max(dp[n-1][0], dp[n-1][1])

	// 选1，那么n-1就不能选
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	res = max(res, dp[n-1][0])
	return res
}
