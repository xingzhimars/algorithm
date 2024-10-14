package main

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	ans := make([]int, n+1)
	ans[0] = 1
	return ans
}
