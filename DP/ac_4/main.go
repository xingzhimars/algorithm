package main

import "fmt"
import "bufio"
import "os"

const N = 110

var n, m int

var v, w, s [N]int

var dp [N][N]int

// 二维数组
func solve() int {

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= s[i] && k*v[i] <= j; k++ {
				// 可以看成是k取0
				// dp[i][j] = dp[i-1][j]

				//             if j >= v[i] {
				//                 dp[i][j] = max(dp[i][j], dp[i-1][j-k*v[i]] + k * w[i])
				//             }

				dp[i][j] = max(dp[i][j], dp[i-1][j-k*v[i]]+k*w[i])

			}
		}
	}

	return dp[n][m]

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
		fmt.Fscan(r, &v[i], &w[i], &s[i])
	}

	fmt.Println(solve())
}
