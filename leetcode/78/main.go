package main

var ans [][]int

func subsets(nums []int) [][]int {
	ans = make([][]int, 0)
	dfs(nums, 0, make([]int, 0))
	return ans
}

func dfs(nums []int, index int, path []int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	ans = append(ans, tmp)
	for i := index; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, i+1, path)
		path = path[:len(path)-1]
	}
}
