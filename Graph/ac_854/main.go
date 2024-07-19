package main

import "fmt"
import "bufio"
import "os"

const N = 210

const M = 20010

var d [N][N]int

var n, m, k int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func floyd() {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m, &k)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = 0x3f3f3f3f
			}
		}
	}

	for m > 0 {
		m--
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)

		d[a][b] = min(d[a][b], c)
	}

	floyd()

	for k > 0 {
		k--
		var x, y int
		fmt.Fscan(r, &x, &y)

		if d[x][y] > 0x3f3f3f3f/2 {
			fmt.Println("impossible")
		} else {
			fmt.Println(d[x][y])
		}
	}
}
