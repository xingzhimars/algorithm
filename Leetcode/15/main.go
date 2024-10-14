package main

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j+1] == nums[j] {
					j++
				}
				j++
				for j < k && nums[k-1] == nums[k] {
					k--
				}
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
