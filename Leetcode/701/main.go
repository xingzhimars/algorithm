package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return dfs(root, val)
}

func dfs(root *TreeNode, val int) *TreeNode {
	if root == nil {
		newNode := &TreeNode{
			Val: val,
		}
		return newNode
	}
	if val < root.Val {
		root.Left = dfs(root.Left, val)
	} else if val > root.Val {
		root.Right = dfs(root.Right, val)
	}
	return root
}
