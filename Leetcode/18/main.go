package main

import "sort"

var ans [][]int

func fourSum(nums []int, target int) [][]int {
	ans = make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			p, q := j+1, n-1

			for p < q {
				sum := nums[i] + nums[j] + nums[p] + nums[q]
				if sum < target {
					p++
				} else if sum > target {
					q--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[p], nums[q]})
					for p < q && nums[p+1] == nums[p] {
						p++
					}
					p++
					for p < q && nums[q-1] == nums[q] {
						q--
					}
					q--
				}
			}
		}
	}
	return ans
}
