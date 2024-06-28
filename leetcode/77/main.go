package main

import "fmt"

var ans [][]int

func combine(n int, k int) [][]int {
	ans = make([][]int, 0)
	dfs(1, k, n, make([]int, 0))
	return ans
}

func dfs(index, k, n int, path []int) {
	if k == 0 {
		// 复制path
		tmp := make([]int, len(path))
		copy(tmp, path)

		ans = append(ans, tmp)
		return
	}
	for i := index; i <= n; i++ {
		// 剪枝，还剩下可选择的元素个数 小于 当前的k
		if n-i+1 < k {
			return
		}
		path = append(path, i)
		dfs(i+1, k-1, n, path)
		path = path[:len(path)-1]
	}
}
func main() {
	res := combine(4, 2)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
