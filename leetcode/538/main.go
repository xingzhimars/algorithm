package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 右中左
func convertBST(root *TreeNode) *TreeNode {
	st := make([]*TreeNode, 0)

	var pre *TreeNode

	cur := root
	for len(st) > 0 || cur != nil {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Right
		} else {
			cur = st[len(st)-1]
			st = st[:len(st)-1]
			if pre != nil {
				cur.Val += pre.Val
			}

			pre = cur

			cur = cur.Left
		}
	}
	return root
}
