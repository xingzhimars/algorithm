package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &k)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	res := findK(nums, 0, n-1, k-1)
	fmt.Println(res)
}

func findK(nums []int, start, end, k int) int {
	if start >= end {
		return nums[start]
	}
	j := partition(nums, start, end)
	if j == k {
		return nums[j]
	}
	if j < k {
		// 注意是j+1，下面同理j-1，始终要清楚是在闭区间里的
		return findK(nums, j+1, end, k)
	} else {
		return findK(nums, start, j-1, k)
	}
}

func partition(nums []int, l, r int) int {
	x := nums[l]
	i, j := l, r+1
	for i < j {
		for i < r {
			i++
			if nums[i] >= x {
				break
			}
		}
		for j > l {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 取左端点为基准点时，最后一步需要交换
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
