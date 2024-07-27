package main

import "fmt"
import "bufio"
import "os"

var s string

var st []bool

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &s)

	st = make([]bool, len(s))

	dfs(s, 0, make([]byte, 0))
}

// start 可以理解成第几个位置，在第几个位置上填写字符
func dfs(s string, start int, path []byte) {
	if start == len(s) {
		for i := 0; i < len(path); i++ {
			fmt.Printf("%c", path[i])
		}
		fmt.Println()
		return
	}
	for i := 0; i < len(s); i++ {
		if !st[i] {
			st[i] = true
			path = append(path, s[i])
			dfs(s, start+1, path)
			path = path[:len(path)-1]
			st[i] = false
		}
	}
}
