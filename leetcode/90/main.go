package main

import "sort"

var ans [][]int

func subsetsWithDup(nums []int) [][]int {
	ans = make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, 0, make([]int, 0))
	return ans
}

func dfs(nums []int, index int, path []int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	ans = append(ans, tmp)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, i+1, path)
		path = path[:len(path)-1]
	}
}
