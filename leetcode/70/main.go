package main

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	f1 := 1
	f2 := 2
	f3 := 0
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}
