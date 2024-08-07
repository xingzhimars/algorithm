package main

import "fmt"
import "bufio"
import "os"

const N = 1010

var n int

var a [N]int

var dp [N]int

// 如果想要记录最长序列的话，在dp中要想记录过程，就是将转移方程记录一下
var g [N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	// dp[i]:以i结尾，严格单调递增的子序列的最长长度
	for i := 0; i < n; i++ {
		g[i] = 0
		for j := 0; j <= i; j++ {
			if a[i] > a[j] {
				// dp[i] = max(dp[i], dp[j] + 1)
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					// i是从j过来的
					g[i] = j
				}
			}
		}
	}
	// 最长子序列不一定是以n-1结尾的
	// res := 0
	// for i := 0; i < n; i++ {
	//     res = max(res, dp[i])
	// }
	// fmt.Println(res)

	k := 0
	for i := 0; i < n; i++ {
		if dp[k] < dp[i] {
			k = i
		}
	}
	fmt.Println(dp[k])

	for i, len := 0, dp[k]; i < len; i++ {
		fmt.Printf("%d ", a[k])
		k = g[k]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
