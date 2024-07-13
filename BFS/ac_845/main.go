package main

import "fmt"

const N = 1e5

// 队列
var queue []string

// map，记录状态和变换次数
var dist map[string]int

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

var end = "12345678x"

func find(str string, s rune) int {
	for i, v := range str {
		if v == s {
			return i
		}
	}
	return -1
}

func swap(str string, i, j int) string {
	bs := []byte(str)
	bs[i], bs[j] = bs[j], bs[i]
	return string(bs)
}

func bfs(state string) int {
	queue = append(queue, state)
	for len(queue) > 0 {
		// st := queue[hh]
		// hh++
		st := queue[0]
		queue = queue[1:]

		distance := dist[st]
		if st == end {
			return distance
		}

		// 先找到当前状态中x的位置
		t := find(st, 'x')
		// 将t转成二维坐标
		x, y := t/3, t%3

		// 搜索周围状态
		for i := 0; i < 4; i++ {
			// 新坐标
			a, b := x+dx[i], y+dy[i]
			// fmt.Printf("%d %d\n", a, b)

			if a >= 0 && a < 3 && b >= 0 && b < 3 {
				// 交换(a, b)与(x, y)的元素，生成新状态，将新状态压入队列
				// 直接在一维中交换
				// st[t], st[a*3+b] = st[a*3+b], st[t]
				st = swap(st, t, a*3+b)

				// 查看是否存在新状态的变换次数，如果变换次数为0，表示之前没有变换过，可以使用
				// 如果变换次数不为0，说明之前使用过，没必要这次使用了，因为肯定不是最少的变换次数
				if _, ok := dist[st]; !ok {
					// 更新新状态的变化次数
					dist[st] = distance + 1

					// 压入队列
					// tt++
					// queue[tt] = st
					queue = append(queue, st)
				}

				// 恢复，因为还要处理其他方向的，而且这个是在当前状态上修改的
				// 一般来说bfs不需要恢复状态，这里的恢复是与具体题目有关
				//  st[t], st[a*3+b] = st[a*3+b], st[t]
				st = swap(st, t, a*3+b)
			}
		}
	}
	return -1
}

func main() {
	var start string
	for i := 0; i < 9; i++ {
		var c string
		fmt.Scan(&c)
		start += c
	}

	dist = make(map[string]int)
	queue = make([]string, 0)

	res := bfs(start)
	fmt.Println(res)
}
