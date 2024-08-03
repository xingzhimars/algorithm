package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// dp[0]：当前节点不偷，那么其左右孩子节点可偷
// dp[1]：当前节点偷，那么其左右孩子节点不可偷
func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	dp := make([]int, 2)
	l := dfs(node.Left)
	r := dfs(node.Right)

	// dp[0] = max(l[0] + l[1], r[0] + r[1])
	// 左孩子中最大值 + 右孩子中最大值
	dp[0] = max(l[0], l[1]) + max(r[0], r[1])
	dp[1] = node.Val + l[0] + r[0]

	return dp
}
