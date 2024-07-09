package main

import "math"

func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
		// 判断res是否越界
		if res >= math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}
	if flag == -1 {
		res = -res
	}
	return res
}
