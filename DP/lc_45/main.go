package main

// 思路：定义f[i]表示跳到第i个位置所需要的最少步数
// 假设 0 <= j < i，从第j个位置，可以跳跃一次到达第i个位置，那么 f[i] = f[j] + 1
// 这个j尽可能的离i远，这样肯定步数最少
// 参考：https://leetcode.cn/problems/jump-game-ii/solutions/775090/gong-shui-san-xie-xiu-gai-shu-ju-fan-wei-wylq/
func jump(nums []int) int {
	n := len(nums)
	f := make([]int, n)

	j := 0
	for i := 1; i < n; i++ {
		for j < i && j+nums[j] < i {
			j++
		}
		f[i] = f[j] + 1
	}
	return f[n-1]
}
