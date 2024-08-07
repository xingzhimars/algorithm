package main

import "fmt"
import "bufio"
import "os"

const N = 1010
const M = 12

var n, m int

var strs [N]string

var a, b [N]byte

var dp [M][M]int

func editDistance(s1, s2 string) int {
	n, m := len(s1), len(s2)

	for i := 1; i <= n; i++ {
		a[i] = s1[i-1]
	}
	for i := 1; i <= m; i++ {
		b[i] = s2[i-1]
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

			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1

			if a[i] == b[j] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &strs[i])
	}

	for m > 0 {
		m--
		var s string
		var limit int
		fmt.Fscan(r, &s, &limit)
		cnt := 0
		for i := 0; i < n; i++ {
			if editDistance(strs[i], s) <= limit {
				cnt++
			}
		}
		fmt.Println(cnt)
	}

}
