package main

const N = 20

var ans [][]string
var matrix [N][N]byte
var cols, dia, udia [N]bool

func solveNQueens(n int) [][]string {
	ans = make([][]string, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = '.'
		}
	}
	dfs(0, n)
	return ans
}

// u：行
func dfs(u, n int) {
	if u == n {
		res := make([]string, 0)
		for i := 0; i < n; i++ {
			path := make([]byte, n)
			for j := 0; j < n; j++ {
				path[j] = matrix[i][j]
			}
			res = append(res, string(path))
		}
		ans = append(ans, res)
		return
	}

	// u行j列，即matrix[u][j]
	for j := 0; j < n; j++ {
		// 同一行不同列，这里就不需要判断rows[i]是否等于false，只要是同一行不同列，不存在皇后
		// 当进入dfs(u+1)时，下一行遍历，如果cols[j]等于true，那么在下一行遍历时，遇到cols[j]等于true时，就表示当前列也存在皇后
		// 因此，这样既可以判断出同一行是否有皇后，也能判断出同一列是否有皇后
		if !cols[j] && !dia[u+j] && !udia[u-j+n] {
			matrix[u][j] = 'Q'
			cols[j], dia[u+j], udia[u-j+n] = true, true, true
			dfs(u+1, n)
			matrix[u][j] = '.'
			cols[j], dia[u+j], udia[u-j+n] = false, false, false
		}
	}
}
