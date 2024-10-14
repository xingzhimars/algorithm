package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	idx := 0

	i, j := 0, 0
	for k := 0; k < n*n; k++ {
		// 加入到当前位置
		matrix[i][j] = k + 1

		// 下一个节点，并且判断是否访问过，坐标是否出界
		// 如果符合条件，那么下一个位置就是它，否则需要换方向
		x := i + dx[idx]
		y := j + dy[idx]
		// 未被访问过 && 未出界
		if x >= 0 && x < n && y >= 0 && y < n && matrix[x][y] == 0 { // 由于题目中都是正数，因此这里可以判断是否等于0
			i = x
			j = y
		} else {
			idx = (idx + 1) % 4
			i += dx[idx]
			j += dy[idx]
		}
	}
	return matrix
}

func main() {
	matrix := generateMatrix(3)
	row := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < row; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
