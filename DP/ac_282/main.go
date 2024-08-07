// dp[i][j]：所有从第i堆到第j堆的合并方式中，最小代价
// 假设k是[i, j]这区间中的，最后一次合并，那么左侧为[i, k]，右侧为[k+1, j]
// 那么dp[i][j] = min(dp[i][k] + dp[k+1][j] + s[j] - s[i-1])

package main

import "fmt"
import "bufio"
import "os"

const N = 310

var n int

var a [N]int
var dp [N][N]int

var preSum [N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &a[i])
		preSum[i] = preSum[i-1] + a[i]
	}

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+preSum[j]-preSum[i-1])
			}
		}
	}

	fmt.Println(dp[1][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
