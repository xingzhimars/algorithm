package main

import "fmt"
import "bufio"
import "os"

const N = 510
const M = 1e5 + 10

var h [N]int
var e, ne [M]int
var idx int

var n1, n2, m int

// 女生是否已经匹配
var st [N]bool

// match[i] = j，i号女生匹配的男生是j号
var match [N]int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func clearSt() {
	for i := 0; i < N; i++ {
		st[i] = false
	}
}

// 男生是否匹配成功
func find(x int) bool {
	for i := h[x]; i != -1; i = ne[i] {
		// j号女生
		j := e[i]
		if !st[j] {
			st[j] = true
			// 如果女生没有被选中 或者 此时与j号女生匹配的男生重新匹配成功
			if match[j] == 0 || find(match[j]) {
				match[j] = x
				return true
			}
		}

	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n1, &n2, &m)

	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		add(a, b)
	}

	res := 0
	for i := 1; i <= n1; i++ {
		// 每次匹配男生时，清空st
		clearSt()
		if find(i) {
			// 如果匹配成功
			res++
		}
	}

	fmt.Println(res)

}
