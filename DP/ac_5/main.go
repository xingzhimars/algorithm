package main

import "fmt"
import "bufio"
import "os"

// 1000 * log2000 = 1000 * 11 = 11000
const N = 15000

var n, m int

var v, w [N]int

var dp [N]int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	cnt := 0
	for i := 0; i < n; i++ {
		var a, b, s int
		fmt.Fscan(r, &a, &b, &s)

		k := 1
		for k <= s {
			cnt++
			v[cnt] = a * k
			w[cnt] = b * k
			s -= k
			k *= 2
		}
		if s > 0 {
			cnt++
			v[cnt] = a * s
			w[cnt] = b * s
		}
	}
	// 更新物品种数
	n = cnt

	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			dp[j] = max(dp[j], dp[j-v[i]]+w[i])
		}
	}
	fmt.Println(dp[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
