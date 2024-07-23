package main

var rows, cols int

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

// 以(i, j)为起始点时，岛屿的面积
func dfs(g [][]int, i, j int) int {
	if i < 0 || i >= rows || j < 0 || j >= cols || g[i][j] == 0 {
		return 0
	}
	g[i][j] = 0
	res := 1
	for k := 0; k < 4; k++ {
		ni, nj := i+dx[k], j+dy[k]
		res += dfs(g, ni, nj)
	}
	return res
}

func maxAreaOfIsland(g [][]int) int {
	rows = len(g)
	cols = len(g[0])

	ans := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g[i][j] == 1 {
				ans = max(ans, dfs(g, i, j))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
