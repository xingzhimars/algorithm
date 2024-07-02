package main

import "bufio"
import "fmt"
import "os"

const N = 50010

var n, m int

// d[x]表示x到p[x]的距离
var p, d [N]int

func find(x int) int {
	if p[x] != x {
		t := find(p[x])
		d[x] += d[p[x]]
		p[x] = t

		// 如果先将p[x]指向根节点，那么d[x]、d[p[x]]的数据就错了
		// p[x] = find(p[x])
		// d[x] += d[p[x]]

	}
	return p[x]
}

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &n, &m)

	for i := 1; i <= n; i++ {
		p[i] = i
	}

	var res int

	for m > 0 {
		m--

		var t, x, y int
		fmt.Fscan(rr, &t, &x, &y)

		if x > n || y > n {
			res++
		} else {
			px, py := find(x), find(y)
			if t == 1 {
				// 分两种情况：1. 在同一个集合中； 2. 在不同集合中
				if px == py {
					if (d[x]-d[y])%3 != 0 {
						res++
					}
				} else {
					// 在不同集合中，合并
					p[px] = py
					// 由于他们是同类，需要更新d
					d[px] = d[y] - d[x]

				}
			} else if t == 2 {
				if px == py {
					// 并非x吃y
					if (d[x]-d[y]-1)%3 != 0 {
						res++
					}
				} else {
					p[px] = py
					d[px] = d[y] + 1 - d[x]
				}
			}
		}

	}
	fmt.Println(res)
}
