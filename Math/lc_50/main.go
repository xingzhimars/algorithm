package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return quick(x, n)
}

// 快速幂
func quick(a float64, b int) float64 {
	var ans float64 = 1.0
	for b > 0 {
		if b&1 == 1 {
			ans *= a
		}
		b = b >> 1
		a *= a
	}
	return ans
}

// 快速乘
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
