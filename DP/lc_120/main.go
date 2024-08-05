package main

func minimumTotal(nums [][]int) int {
	// 考虑从下往上
	for i := len(nums) - 2; i >= 0; i-- {
		// 像这类三角形，注意j的范围，无需 j < cols
		for j := 0; j <= i; j++ {
			nums[i][j] = min(nums[i+1][j], nums[i+1][j+1]) + nums[i][j]
		}
	}
	return nums[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
