package main

import "sort"

var ans [][]int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	ans = make([][]int, 0)
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
		dfs(nums, i+1, target-nums[i], path)
		path = path[:len(path)-1]
	}

}
