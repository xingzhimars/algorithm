package main

func maxSubArray(nums []int) int {
	n := len(nums)
	var maxSum = -0x3f3f3f3f
	var sum int
	for i := 0; i < n; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
