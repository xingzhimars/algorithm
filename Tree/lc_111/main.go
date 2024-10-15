package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

const N = 1e5 + 10

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0

	queue := make([]*TreeNode, N)
	tt := -1
	hh := 0

	tt++
	queue[tt] = root

	lastNode := root
	var nLastNode *TreeNode

	for hh <= tt {
		cur := queue[hh]
		hh++

		if cur.Left == nil && cur.Right == nil {
			// ans++
			break
		}
		if cur.Left != nil {
			tt++
			queue[tt] = cur.Left
			nLastNode = cur.Left
		}
		if cur.Right != nil {
			tt++
			queue[tt] = cur.Right
			nLastNode = cur.Right
		}
		if cur == lastNode {
			ans++
			lastNode = nLastNode
		}
	}

	return ans + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
