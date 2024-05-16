package main

/**
422. 校门外的树
某校大门外长度为 L 的马路上有一排树，每两棵相邻的树之间的间隔都是 1 米。
我们可以把马路看成一个数轴，马路的一端在数轴 0 的位置，另一端在 L 的位置；数轴上的每个整数点，即 0，1，2，……，L，都种有一棵树。
由于马路上有一些区域要用来建地铁。
这些区域用它们在数轴上的起始点和终止点表示。
已知任一区域的起始点和终止点的坐标都是整数，区域之间可能有重合的部分。
现在要把这些区域中的树（包括区域端点处的两棵树）移走。
你的任务是计算将这些树都移走后，马路上还有多少棵树。

输入格式
输入文件的第一行有两个整数 L 和 M，L 代表马路的长度，M 代表区域的数目，L 和 M 之间用一个空格隔开。
接下来的 M 行每行包含两个不同的整数，用一个空格隔开，表示一个区域的起始点和终止点的坐标。

输出格式
输出文件包括一行，这一行只包含一个整数，表示马路上剩余的树的数目。
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var L, M int

var x, y int
var s []*Pair = make([]*Pair, 0)
var ans []*Pair = make([]*Pair, 0)

type Pair struct {
	l, r int
}

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &L, &M)

	for i := 0; i < M; i++ {
		fmt.Fscan(rr, &x, &y)
		s = append(s, &Pair{
			l: x,
			r: y,
		})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].l < s[j].l
	})

	ans = append(ans, s[0])

	for i, length := 1, len(s); i < length; i++ {
		pre := ans[len(ans)-1]

		if pre.r < s[i].l {
			ans = append(ans, s[i])
		} else {
			pre.r = max(pre.r, s[i].r)
		}
	}
	var sum int
	for i, length := 0, len(ans); i < length; i++ {
		sum += ans[i].r - ans[i].l + 1
	}
	// 标记为0的地方需要种树
	fmt.Println(L + 1 - sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
