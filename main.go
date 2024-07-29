package main

func generateMatrix(n int) [][]int {
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	d := 0

	i, j := 0, 0
	for k := 1; k <= n*n; k++ {
		nums[i][j] = k

		ni, nj := i+dx[d], j+dy[d]
		if ni < 0 || ni >= n || nj < 0 || nj >= n || nums[ni][nj] != 0 {
			d = (d + 1) % 4
			ni = i + dx[d]
			nj = j + dy[d]
		}
		i = ni
		j = nj
	}
	return nums
}

func main() {
	_ = generateMatrix(3)
}
