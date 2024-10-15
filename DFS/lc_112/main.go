package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	var dfs func(*TreeNode, int) bool
	dfs = func(root *TreeNode, target int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil && target == root.Val {
			return true
		}
		res1 := dfs(root.Left, target-root.Val)
		res2 := dfs(root.Right, target-root.Val)
		return res1 || res2
	}
	return dfs(root, targetSum)
}
