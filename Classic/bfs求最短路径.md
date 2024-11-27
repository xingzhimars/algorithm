### bfs求最短路径      
bfs一般就是使用队列处理，只有边权都为1的情况下，才会使用bfs求最短路径  




题目  
给定一个 n×m 的二维整数数组，用来表示一个迷宫，数组中只包含 0 或 1 ，其中 0 表示可以走的路，1 表示不可通过的墙壁。

最初，有一个人位于左上角 (1,1)处，已知该人每次可以向上、下、左、右任意一个方向移动一个位置。 
请问，该人从左上角移动至右下角 (n,m)处，至少需要移动多少次。 
数据保证 (1,1)处和 (n,m)处的数字为 0 ，且一定至少存在一条通路。

输入格式   
第一行包含两个整数 n 和 m 。 接下来 n 行，每行包含 m 个整数（0 或 1 ），表示完整的二维数组迷宫。

输出格式   
输出一个整数，表示从左上角移动至右下角的最少移动次数。

数据范围   
1≤n,m≤100

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 110

type Pair struct {
	i, j int
}

var n, m int

var g [N][N]int
var d [N][N]int // 距离数组

var queue [N * N]*Pair

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &g[i][j])
		}
	}

	// 初始化d
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			d[i][j] = -1
		}
	}

	d[0][0] = 0
	bfs()
	fmt.Println(d[n-1][m-1])
}

func bfs() {
	head, tail := 0, -1
	tail++
	queue[tail] = &Pair{
		i: 0,
		j: 0,
	}
	for head <= tail {
		cur := queue[head]
		head++

		for k := 0; k < 4; k++ {
			x := cur.i + dx[k]
			y := cur.j + dy[k]

			// d[x][y] == -1 表示没走过
			if x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && d[x][y] == -1 {
				d[x][y] = d[cur.i][cur.j] + 1

				// 满足条件，加入到队列
				tail++
				queue[tail] = &Pair{
					i: x,
					j: y,
				}
			}
		}
	}
}
```

