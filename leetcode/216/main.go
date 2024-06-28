package main

var ans [][]int

func combinationSum3(k int, n int) [][]int {
	ans = make([][]int, 0)
	dfs(1, n, k, make([]int, 0))
	return ans
}

func dfs(idx, target, k int, path []int) {
	if k <= 0 {
		if target == 0 && k == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		return
	}
	for i := idx; i <= 9; i++ {
		path = append(path, i)
		dfs(i+1, target-i, k-1, path)
		path = path[:len(path)-1]
	}
}
