package main

import "math"

// 快速乘
func quickMul(a, b int) int {
	ans := 0
	for b > 0 {
		if b&1 == 1 {
			ans += a
		}
		b = b >> 1
		a += a
	}
	return ans
}

func divide(x int, y int) int {
	flag := true
	if x > 0 && y < 0 {
		flag = false
		y = -y
	}
	if x < 0 && y > 0 {
		flag = false
		x = -x
	}
	if x < 0 && y < 0 {
		x = -x
		y = -y
	}

	l, r := 0, x
	for l < r {
		mid := (l + r + 1) >> 1
		if quickMul(y, mid) <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	ans := r
	if !flag {
		ans = -ans
	}
	if ans > 0 && ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < 0 && ans < math.MinInt32 {
		return math.MinInt32
	}
	return ans
}
