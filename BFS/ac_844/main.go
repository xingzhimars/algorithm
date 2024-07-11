package main

// 题目：走迷宫
// 给定一个n*m的二维数组，用来表示迷宫，数组中只包含0或者1，其中0表示可以走的路，1表示不可以通过的墙壁
// 从左上角(0, 0)走到右下角(n-1, m-1)，至少需要移动多少次

import "fmt"
import "bufio"
import "os"

const N = 110

var n, m int

// 矩阵数组
var g [N][N]int

// 距离数组，记录点到起点的距离，-1表示没有走过
var d [N][N]int

// 队列
var queue [N * N]*Pair
var head, tail int

type Pair struct {
	x, y int
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func bfs() {

	head, tail := 0, -1

	tail++
	queue[tail] = &Pair{
		x: 0,
		y: 0,
	}

	// 这个不能少，(0,0)到起始点的距离是0
	d[0][0] = 0

	for head <= tail {
		t := queue[head]
		head++

		// 上下左右四个方向搜索
		for i := 0; i < 4; i++ {
			x := t.x + dx[i]
			y := t.y + dy[i]

			// 判断是否满足条件
			if x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && d[x][y] == -1 {
				d[x][y] = d[t.x][t.y] + 1

				tail++
				queue[tail] = &Pair{
					x: x,
					y: y,
				}
			}
		}
	}

}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &g[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			d[i][j] = -1
		}
	}

	bfs()

	fmt.Println(d[n-1][m-1])
}
