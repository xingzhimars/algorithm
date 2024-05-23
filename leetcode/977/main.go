package main

// 题目：给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 思路：双指针
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	k := n - 1
	start, end := 0, n-1
	for start <= end {
		v1 := nums[start] * nums[start]
		v2 := nums[end] * nums[end]
		if v1 > v2 {
			ans[k] = v1
			start++
		} else {
			ans[k] = v2
			end--
		}
		k--
	}
	return ans
}
