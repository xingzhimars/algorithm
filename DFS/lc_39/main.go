package main

import (
	"sort"
)

var ans [][]int

// 思路：找不同组合的，需要先将数组排序，去重
func combinationSum(candidates []int, target int) [][]int {
	ans = make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	dfs(candidates, 0, target, make([]int, 0))
	return ans
}

func dfs(nums []int, idx, target int, path []int) {
	if target <= 0 {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		return
	}

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, i, target-nums[i], path)
		path = path[:len(path)-1]
	}

}
