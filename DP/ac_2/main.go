package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var v, w [N]int

var n, m int

var dp [N]int

// 定义：dp[i][j]表示使用前i个物品，总体积不超过j的最大价值
// 初始化：
// 状态转移：
// 选择i：dp[i][j] = dp[i-1][j-v[i]] + w[i]
// 不选择i：dp[i][j] = dp[i-1][j]

// 一维数组只是对二维数组的等价变换
func solve() int {
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			dp[j] = max(dp[j], dp[j-v[i]]+w[i])
		}
	}
	return dp[m]
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

	fmt.Println(solve())
}
