package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp == 0 {
				res := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, res)
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if tmp < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
