package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

var n int

var son [N][26]int
var cnt [N]int
var idx int

func insert(s string) {
	// 指向根节点
	p := 0
	for i := 0; i < len(s); i++ {
		u := s[i] - 'a'
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	// 以p结尾的字符串个数
	cnt[p]++
}

func query(s string) int {
	p := 0
	for i := 0; i < len(s); i++ {
		u := s[i] - 'a'
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	for n > 0 {
		n--

		var op, s string
		fmt.Fscan(r, &op, &s)

		if op == "I" {
			insert(s)
		} else if op == "Q" {
			res := query(s)
			fmt.Println(res)
		}
	}
}
