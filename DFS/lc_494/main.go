package main

func findTargetSumWays(nums []int, target int) int {
	cnt := 0

	var dfs func([]int, int, int)
	dfs = func(nums []int, start, target int) {
		if start == len(nums) {
			if 0 == target {
				cnt++
			}
			return
		}
		// 不需要在当前位置还遍历后续元素，只需要关心当前位置是正号，还是负号
		// for i := start; i < len(nums); i++ {
		// 	dfs(nums, start+1, target - nums[i])
		// 	dfs(nums, start+1, target + nums[i])
		// }
		dfs(nums, start+1, target-nums[start])
		dfs(nums, start+1, target+nums[start])
	}
	dfs(nums, 0, target)
	return cnt
}
