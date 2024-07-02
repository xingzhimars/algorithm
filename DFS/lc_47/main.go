package main

import "sort"

var ans [][]int
var used []bool

// 思路：包含重复的元素，且要求不重复的结果 -> 先排序，排序去重
func permuteUnique(nums []int) [][]int {
	ans = make([][]int, 0)
	used = make([]bool, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, 0, make([]int, 0))
	return ans
}

func dfs(nums []int, index int, path []int) {
	if index == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 一定要加上!used[i-1]。
		// 假设[1,1,2]，一定是nums[i-1]没有使用过，才跳过，比如：第一个1使用过，第二个1完全是可以用的！这个不重复
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		dfs(nums, index+1, path)
		path = path[:len(path)-1]
		used[i] = false
	}
}
