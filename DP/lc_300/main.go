package main

// dp[i]: 以下标i为结尾的元素，最长递增子序列的长度
// 不能是定义成前i个元素，因为需要判断nums[i] > nums[j]
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+10)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
