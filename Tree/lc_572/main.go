package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSubtreeErr(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot == nil {
		return root == nil
	}
	if root.Val != subRoot.Val {
		return isSubtreeErr(root.Left, subRoot) || isSubtreeErr(root.Right, subRoot)
	}
	// 问题在这里，根节点的值相同，此时左右子节点不同，此时subRoot都已经遍历到叶子节点了，无法再从头开始遍历
	return isSubtreeErr(root.Left, subRoot.Left) && isSubtreeErr(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isSame(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil {
		return root1 == nil
	}
	return root1.Val == root2.Val && isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}
