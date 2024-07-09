package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			p := n - 1
			for k < p {
				sum := nums[i] + nums[j] + nums[k] + nums[p]
				if sum == target {
					tmp := []int{nums[i], nums[j], nums[k], nums[p]}
					ans = append(ans, tmp)
					k++
					for k < p && nums[k] == nums[k-1] {
						k++
					}
					p--
					for k < p && nums[p] == nums[p+1] {
						p--
					}
				} else if sum < target {
					k++
				} else {
					p--
				}
			}
		}
	}
	return ans
}
