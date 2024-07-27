package main

// 前缀异或
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	m := len(queries)

	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] ^ arr[i-1]
	}

	ans := make([]int, m)
	for i := 0; i < m; i++ {
		l, r := queries[i][0]+1, queries[i][1]+1
		ans[i] = preSum[r] ^ preSum[l-1]
	}
	return ans
}
