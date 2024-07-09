package main

// 双指针
func maxArea(height []int) int {
	n := len(height)

	maxArea := 0
	start, end := 0, n-1
	for start < end {
		y := min(height[start], height[end])
		x := end - start
		maxArea = max(maxArea, x*y)
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
