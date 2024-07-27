package main

func maxProfit(prices []int) int {
	n := len(prices)
	min := 0x3f3f3f3f
	res := 0
	for i := 0; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		res = max(res, prices[i]-min)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
