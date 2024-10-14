package main

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1

	res := 0

	for l < r {
		y := min(height[l], height[r])
		x := r - l
		res = max(res, x*y)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
