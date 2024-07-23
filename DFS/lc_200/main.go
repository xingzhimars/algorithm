package main

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

var rows, cols int

func numIslands(g [][]byte) int {
	rows = len(g)
	cols = len(g[0])

	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g[i][j] == '1' {
				dfsI(g, i, j)
				res++
			}
		}
	}
	return res
}

// 每次dfs，都将附近陆地变成水，那么执行多少次dfs，就表示有多少个岛屿
func dfs(g [][]byte, i, j int) {
	g[i][j] = '0'

	for k := 0; k < 4; k++ {
		ni := i + dx[k]
		nj := j + dy[k]
		if ni >= 0 && ni < rows && nj >= 0 && nj < cols && g[ni][nj] == '1' {
			dfs(g, ni, nj)
		}
	}
}

// 两种dfs写法都可以，就是一个是下一次递归判断条件，一个是当前递归就判断条件
func dfsI(g [][]byte, i, j int) {
	if i < 0 || i >= rows || j < 0 || j >= cols || g[i][j] == '0' {
		return
	}

	g[i][j] = '0'

	for k := 0; k < 4; k++ {
		ni, nj := i+dx[k], j+dy[k]
		dfsI(g, ni, nj)
	}
}
