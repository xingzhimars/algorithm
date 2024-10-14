package main

// 双指针，国旗问题
func sortColors(nums []int) {
	n := len(nums)
	start := 0   // start指向最后一个红色
	end := n - 1 // end指向第一个蓝色

	i := 0

	// i < n 是错误的
	for i <= end {
		if nums[i] == 0 {
			nums[i], nums[start] = nums[start], nums[i]
			start++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		} else {
			i++
		}
	}
}
