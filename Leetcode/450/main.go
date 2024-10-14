package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	return dfs(root, key)
}

func dfs(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 1. 叶子节点
		if root.Left == nil && root.Right == nil {
			// 删除叶子节点
			return nil
		}
		// 2. 左不为空，右为空
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// 3. 左为空，右不为空
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		// 4. 左不为空，右不为空
		if root.Left != nil && root.Right != nil {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	// 单层递归的逻辑
	if root.Val < key {
		// 返回值是以root.Right为根的子树，再删除节点后，返回的新子树的根
		root.Right = dfs(root.Right, key)
	}
	if root.Val > key {
		root.Left = dfs(root.Left, key)
	}
	return root
}
