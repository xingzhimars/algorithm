package main

import "fmt"
import "bufio"
import "os"

const N = 20

var g [N][N]byte

var n, m int

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

var st [26]int

func add(row int, s string) {
	for i := 0; i < len(s); i++ {
		g[row][i] = s[i]
	}
}

var ans int

// 过滤条件写在循环里面，本次循环的点是ok的
func dfs(i, j, cnt int) {
	ans = max(ans, cnt)
	for k := 0; k < 4; k++ {
		ni, nj := i+dx[k], j+dy[k]
		if ni >= 0 && ni < n && nj >= 0 && nj < m && st[g[ni][nj]-'A'] == 0 {
			st[g[ni][nj]-'A']++
			dfs(ni, nj, cnt+1)
			st[g[ni][nj]-'A']--
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < n; i++ {
		var a string
		fmt.Fscan(r, &a)
		add(i, a)
	}

	st[g[0][0]-'A']++
	dfs(0, 0, 1)
	fmt.Println(ans)
}
