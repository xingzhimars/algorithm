package main

import "fmt"
import "bufio"
import "os"

const N = 110

var n, m int

var v, w [N][N]int
var s [N]int
var dp [N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 1; i <= n; i++ {

		fmt.Fscan(r, &s[i])

		for j := 1; j <= s[i]; j++ {
			fmt.Fscan(r, &v[i][j], &w[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := m; j >= 0; j-- {
			// 第i组的第k个物品
			for k := 0; k <= s[i]; k++ {
				if j >= v[i][k] {
					dp[j] = max(dp[j], dp[j-v[i][k]]+w[i][k])
				}
			}
		}
	}
	fmt.Println(dp[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
