package main

import "fmt"
import "bufio"
import "os"

const N = 510

var n int

var a, dp [N][N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fscan(r, &a[i][j])
		}
	}

	// 自顶向下走，需要判断边界
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = max(dp[i+1][j], dp[i+1][j+1]) + a[i][j]
		}
	}

	fmt.Println(dp[0][0])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
