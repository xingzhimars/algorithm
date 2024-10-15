package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var ans bool = true
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := dfs(root.Left)
		rh := dfs(root.Right)
		if abs(lh-rh) > 1 {
			ans = false
		}
		return max(lh, rh) + 1
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
