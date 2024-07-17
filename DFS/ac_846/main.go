package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

// 两条边
const M = N * 2

// 邻接表实现图
// var h, e, ne [N]int
var h [N]int
var e, ne [M]int
var idx int

// 一般情况下，图的遍历，点只会遍历一次
var visit [N]bool

var n int

// 求最小值，初始化时就需要赋值个较大值
var ans int = N

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &n)

	// 初始化 h
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		// 构造无向图
		// 无向图是特殊的有向图，在构造的过程中，连两条边即可
		add(a, b)
		add(b, a)
	}
	dfs(1)
	fmt.Println(ans)
}

// 添加一条从a指向b的边
func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx // 新插入的节点变成头节点
	idx++
}

// 以u为根的子树中点的个数
func dfs(u int) int {
	visit[u] = true

	// sum表示以u为根的子树中点的个数，res表示连通块中点数的最大值
	sum, res := 1, 0
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if !visit[j] {
			childSize := dfs(j)
			// 以u为根，childSize其实就是去掉u之后的一个连通块的个数（画图就能理解）
			res = max(res, childSize)
			sum += childSize
		}
	}
	// 连通块还有一部分，就是除去u为根的子树后的部分
	res = max(res, n-sum)
	ans = min(ans, res)
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
