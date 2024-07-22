// 题目：圆桌座位
// N 个人围坐一圈，有 M 对朋友关系。
// 第 i 对朋友关系是指，编号是 ai 的人和编号是 bi 的人是朋友。
// 现在要给他们安排座位，要求所有相邻的人不能是朋友。
// 问共有多少种方案？ 如果两个方案只有旋转角度不同，则我们将其视为一种方案。

package main

import "fmt"
import "bufio"
import "os"

const N = 20
const M = 200

// 是否为朋友
var g [N][N]bool

// i号同志是否用过
var st [N]bool

// 下标i位置上的人是几号同志
var pos [N]int

var n, m int

// u第几个位置，返回方案数
func dfs(u int) int {
	// 一次遍历完成
	if u == n {
		// n-1位置上的同志与0位置上的同志是否为朋友，如果是非朋友，那么当前pos数组顺序是一种方案
		if g[pos[n-1]][pos[0]] {
			return 0
		}
		return 1
	}
	res := 0
	// 遍历每个同志
	for i := 1; i <= n; i++ {
		// 如果i号没用过，且i号与前一个位置的人非朋友，将i号放在u位置
		if !st[i] && !g[i][pos[u-1]] {
			pos[u] = i
			st[i] = true
			res += dfs(u + 1)
			st[i] = false
			// pos会覆盖，无需还原
		}
	}
	return res
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)

		g[a][b] = true
		g[b][a] = true
	}

	// 下标为0的位置上是1号同志；1号同志已经用过
	pos[0] = 1
	st[1] = true

	// 开始处理下标为1的位置
	t := dfs(1)

	fmt.Println(t)
}
