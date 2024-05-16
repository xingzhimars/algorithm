package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：由于是二叉搜索树，可以利用其特性
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left, right := p.Val, q.Val
	if left > right {
		left, right = right, left
	}
	if root.Val < left {
		return lowestCommonAncestor235(root.Right, p, q)
	} else if root.Val > right {
		return lowestCommonAncestor235(root.Left, p, q)
	}
	return root
}
