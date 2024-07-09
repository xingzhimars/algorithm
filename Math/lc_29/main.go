package main

import "math"

// 两数相除，不能使用乘法、除法和取余
// 思路：二分法 && 快速乘
func divide(dividend int, divisor int) int {
	x, y := dividend, divisor
	var flag = true
	if x < 0 && y < 0 {
		flag = true
		x = -x
		y = -y
	}
	if x < 0 && y > 0 {
		x = -x
		flag = false
	}
	if x > 0 && y < 0 {
		y = -y
		flag = false
	}

	l, r := 0, x
	for l < r {
		mid := (l + r + 1) >> 1
		if quickMul(mid, y) <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	res := l
	if !flag {
		res = -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	return res
}

// 快速乘：将乘法改成加法的形式
func quickMul(a, b int) int {
	var ans int
	for b > 0 {
		if b&1 == 1 {
			ans += a
		}
		b = b >> 1
		a += a
	}
	return ans
}
