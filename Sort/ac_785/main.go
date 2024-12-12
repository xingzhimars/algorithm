package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	quickSort(nums, 0, n-1)

	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	j := partition(nums, start, end)
	quickSort(nums, start, j)
	quickSort(nums, j+1, end)
}

func partition(nums []int, start, end int) int {
	x := nums[start+(end-start)/2]
	i, j := start-1, end+1

	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
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
	return j
}
