// dp[i][j]：A字符串的前i个字符，B字符串的前j个字符，它们的最长子序列
// 1. A[i]、B[j] 都不选  dp[i-1][j-1]
// 2. A[i]、B[j] 都选 dp[i-1][j-1] + 1
// 3. A[i]不选，B[j]选  不是dp[i-1][j]，因为dp[i-1][j]包含着B[j]不选，但是也包含着B[j]选，
//    由于是求最大值，dp[i][j]包含dp[i-1][j]，dp[i-1][j]包含着"A[i]不选，B[j]选"
// 4. A[i]选，B[j]不选  同上

package main

import "fmt"
import "bufio"
import "os"

const N = 1010

var n, m int

var dp [N][N]int
var a, b [N]byte

// dp问题中，如果出现了i-1等，下标最好从1开始
func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	var x, y string
	fmt.Fscan(r, &x, &y)

	for i := 1; i <= n; i++ {
		a[i] = x[i-1]
	}

	for j := 1; j <= m; j++ {
		b[j] = y[j-1]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			dp[i][j] = max(dp[i-1][j], dp[i][j-1])

			if a[i] == b[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	fmt.Println(dp[n][m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
