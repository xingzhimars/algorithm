// dp[i][j]：A前i个字符变换成B前j个字符，最少需要的操作次数
// 状态转移方程：
// 1. 若A[i] == B[j]，不需要操作
// 2. 若A[i] != B[j]
//    2.1 A[i-1] == B[j-1]，则将A[i]替换成B[j]，操作次数 + 1，dp[i-1][j-1] + 1

// 3. A[i-1] == B[j]，则删除A[i]，操作次数 + 1，dp[i-1][j] + 1
// 4. A[i] == B[j-1]，则在A[i]后面插入一个字符B[j]，操作次数 + 1，dp[i][j-1] + 1

package main

import "fmt"
import "bufio"
import "os"

const N = 1010

var n, m int

var a, b [N]byte

var dp [N][N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	var x string
	fmt.Fscan(r, &x)

	for i := 1; i <= n; i++ {
		a[i] = x[i-1]
		// fmt.Printf("%c", a[i])
	}

	fmt.Fscan(r, &m)

	var y string
	fmt.Fscan(r, &y)

	// fmt.Println()

	for i := 1; i <= m; i++ {
		b[i] = y[i-1]
		// fmt.Printf("%c", b[i])
	}

	for i := 1; i <= n; i++ {
		// 删除字符
		dp[i][0] = i
	}

	for i := 1; i <= m; i++ {
		// 插入字符
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			// if a[i] != b[j] {
			//     dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			// }
			// 先计算3，4
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1

			if a[i] == b[j] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	fmt.Println(dp[n][m])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
