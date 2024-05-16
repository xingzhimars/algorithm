package main

/**
1343. 挤牛奶
每天早上 5 点，三名农夫去牛场给奶牛们挤奶。
现在从 5 点开始按秒计时，第一名农夫在第 300 秒开始给牛挤奶，并在第 1000 秒停止挤奶。
第二名农夫在第 700 秒开始给牛挤奶，并在第 1200 秒停止挤奶。
第三名农夫在第 1500 秒开始给牛挤奶，并在第 2100 秒停止挤奶。

从开始挤奶到挤奶完全结束，这一期间，至少存在一名农夫正在挤奶的连续时间段的长度最长为 900 秒（第 300 秒至第 1200 秒），
完全没有任何农夫在挤奶的连续时间段的长度最长为 300 秒（第 1200 秒至第 1500 秒）。

现在给你 N 名农夫挤 N 头奶牛的工作时间表，请你求出：
至少存在一名农夫正在挤奶的连续时间段的最长长度。
没有任何农夫在挤奶的连续时间段的最长长度。

注意：本题中给出的所有时间均为时刻（时间点），因此在本题中挤奶区间 [100，200] 和 [201,300] 中间会有长度为 1 秒的间歇时间。

输入格式
第一行包含整数 N，表示农夫数量。
接下来 N 行，每行包含两个非负整数 l,r，表示农夫挤奶的开始时刻和结束时刻。

输出格式
共一行，包含两个整数，分别表示最长连续挤奶时间以及最长连续无人挤奶时间。
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var N int
var x, y int
var s = make([]*Pair, 0)
var ans = make([]*Pair, 0)

type Pair struct {
	l, r int
}

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &N)

	for i := 0; i < N; i++ {
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
	for i, length := 0, len(s); i < length; i++ {
		pre := ans[len(ans)-1]
		if pre.r < s[i].l {
			ans = append(ans, s[i])
		} else {
			pre.r = max(pre.r, s[i].r)
		}
	}

	var p1, p2 int
	for i, length := 0, len(ans); i < length; i++ {
		p1 = max(p1, ans[i].r-ans[i].l)
		if i > 0 {
			p2 = max(p2, ans[i].l-ans[i-1].r)
		}
	}
	fmt.Printf("%d %d\n", p1, p2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
