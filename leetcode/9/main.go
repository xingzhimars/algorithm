package main

func isPalindrome(x int) bool {
	tmp := x
	if tmp < 0 {
		return false
	}
	res := 0
	for tmp > 0 {
		res = res*10 + tmp%10
		tmp /= 10
	}
	return res == x
}
