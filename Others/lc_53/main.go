package main

import "math"

func maxSubArray(nums []int) int {
	n := len(nums)
	maxSum := math.MinInt32
	sum := 0
	for i := 0; i < n; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
