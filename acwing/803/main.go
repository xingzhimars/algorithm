package main

/**
区间合并：
给定 n 个区间 [li,ri]，要求合并所有有交集的区间。注意如果在端点处相交，也算有交集。
输出合并完成后的区间个数。
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n int
var x, y int

type Pair struct {
	l, r int
}

var s []*Pair = make([]*Pair, 0)
var ans []*Pair = make([]*Pair, 0)

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &n)

	for i := 0; i < n; i++ {
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
	fmt.Println(len(ans))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
