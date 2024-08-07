package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var n, m int

var v, w [N]int

var dp [N][N]int

var dp2 [N]int

func solve() int {

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {

			// 不选择i，其实下面k等于0，等价于这个，可以合并
			dp[i][j] = dp[i-1][j]

			if j >= v[i] {
				for k := 0; k*v[i] <= j; k++ {
					// 选择i，选择多少呢?
					dp[i][j] = max(dp[i][j], dp[i-1][j-k*v[i]]+k*w[i])
				}
			}

		}
	}
	return dp[n][m]
}

// 一维数组
// dp[i][j] = max(dp[i-1][j], dp[i-1][j-v[i]] + w[i], dp[i-1][j-2*v[i]] + 2*w[i], dp[i-1][j-3*v[i]] + 3*w[i], ...)
// dp[i][j-v[i]] = max(       dp[i-1][j-v[i]],        dp[i-1][j-2*v[i]] + w[i],   dp[i-1][j-3*v[i]] + 2*w[i])

// 那么 dp[i][j] = max(dp[i-1][j], dp[i][j-v[i]]+w[i])
func solve_2() int {
	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			dp2[j] = max(dp2[j], dp2[j-v[i]]+w[i])
		}
	}
	return dp2[m]
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

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}

	fmt.Println(solve_2())
}
