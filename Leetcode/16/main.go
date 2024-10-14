package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := 0x3f3f3f3f
	for i := 0; i < n; i++ {
		j, k := i+1, n-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if abs(res, target) > abs(sum, target) {
				res = sum
			}

			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
