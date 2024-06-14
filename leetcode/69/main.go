package main

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)>>1

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
