package main

import (
	"fmt"
	"math"
)

// 题目：给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

// 思路：j表示滑动窗口的终止位置，当j滑动的时候，计算sum，如果sum大于等于target时，移动起始位置i，在这个过程中计算minLen
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	i := 0
	sum := 0
	// 求最小值，初始值一定得为最大的
	minLen := math.MaxInt64
	for j := 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			minLen = min(minLen, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if minLen == math.MaxInt64 {
		return 0
	}
	return minLen
}

// 这样写，会超时
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)

	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	minLen := 0x3f3f3f3f

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			sum := preSum[j] - preSum[i-1]
			if sum >= target {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
			}
		}
	}
	if minLen == 0x3f3f3f3f {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
