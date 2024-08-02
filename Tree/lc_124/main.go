package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	var maxLen int = -0x3f3f3f3f
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l < 0 {
			l = 0
		}
		if r < 0 {
			r = 0
		}
		maxLen = max(maxLen, node.Val+l+r)
		return node.Val + max(l, r)
	}

	dfs(root)
	return maxLen
}
