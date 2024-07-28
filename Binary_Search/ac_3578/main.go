package main

import "fmt"
import "bufio"
import "os"
import "sort"

// 思路：二分法
// 假设[L, R]，中间值为mid，假设mid为中位数，那么要想使得mid为中位数，那么数组中后半部分的元素分别减去mid值，
// 如果小于等于k，表示mid是可以成为数组的中位数的，但是要想中位数更大，可以再次将mid值变大一些

const N = 2e5 + 10

var n, k int

var a []int

func check(u int) bool {
	var res int
	for i := n / 2; i < n; i++ {
		if a[i] < u {
			res += u - a[i]
		}
	}
	return res <= k
}

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &n, &k)

	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rr, &a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	// r应该从2e9开始，k最大为1e9，将ai加k次1，就是2e9
	l, r := 0, int(2e9)

	var mid int
	for l < r {
		mid = (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Println(l)
}
