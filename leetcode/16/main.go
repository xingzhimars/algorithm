package main

import "sort"

// 思路：计算sum-target的最小值，在这过程中，记录最小值
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	ans := nums[0] + nums[1] + nums[2]
	for k := 0; k+2 < n; k++ {
		i := k + 1
		j := n - 1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if abs(ans, target) > abs(sum, target) {
				ans = sum
			}
			if sum > target {
				j--
			} else if sum < target {
				i++
			} else {
				return sum
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	tmp := a - b
	if tmp < 0 {
		tmp = -tmp
	}
	return tmp
}
