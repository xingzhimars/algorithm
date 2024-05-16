package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：中序遍历是有序的
func isValidBST(root *TreeNode) bool {
	st := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for len(st) > 0 || cur != nil {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
		} else {
			cur = st[len(st)-1]
			st = st[:len(st)-1]
			if pre != nil {
				if pre.Val >= cur.Val {
					return false
				}
			}
			pre = cur
			cur = cur.Right
		}
	}
	return true
}
