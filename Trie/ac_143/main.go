package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

// M Trie树总节点的个数
const M = 31*N + 10

var son [M][2]int
var idx int

var n int
var a [N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		// 构造trie树
		insert(a[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, query(a[i]))
	}

	fmt.Println(res)

}

func insert(x int) {
	p := 0
	for i := 30; i >= 0; i-- {
		u := (x >> i) & 1
		// 如果不存在，创建新节点
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
}

func query(x int) int {
	p := 0
	res := 0
	for i := 30; i >= 0; i-- {
		u := (x >> i) & 1
		if son[p][1-u] != 0 {
			// 记录res的值
			res += 1 << i
			p = son[p][1-u]
		} else {
			p = son[p][u]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
