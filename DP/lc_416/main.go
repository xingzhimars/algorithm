package main

// 题目：给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 思路：假设元素和为sum，那么就是求是否可以用数组中的元素组成元素之和为sum / 2
// dp[i]：考虑前i个元素，是否可以组成元素之和为sum / 2
// 0-1背包问题，可选可不选
func canPartition(nums []int) bool {
	n := len(nums)

	sum := 0
	maxVal := -0x3f3f3f3f
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if sum%2 != 0 {
		return false
	}
	if maxVal > sum/2 {
		return false
	}

	target := sum / 2

	dp := make([]bool, target+1)

	dp[0] = true

	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}
