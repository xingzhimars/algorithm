package main

var ans [][]int

var used []bool

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	used = make([]bool, len(nums))
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
	// 全排列，所有元素都必须包含，因此这里的i从0开始，而不能从index开始，
	// 从0开始，又会出现一个问题，即同一个元素会被选择多次，因此又需要判断是否使用过
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		// 这里的i不能从i+1开始，这样也是不包含所有元素的，而应该是从index+1开始，
		// index可以理解成当前填空的下标
		dfs(nums, index+1, path)
		path = path[:len(path)-1]
		used[i] = false
	}
}
