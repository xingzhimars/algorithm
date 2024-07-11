package main

func canJump(nums []int) bool {
	n := len(nums)
	// 可跳跃的最远距离
	var maxLen int

	for i := 0; i <= maxLen; i++ {
		if nums[i]+i > maxLen {
			maxLen = nums[i] + i
		}
		if maxLen >= n-1 {
			return true
		}
	}
	return false
}
