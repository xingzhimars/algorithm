package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：中序遍历数组，计算相邻值
func getMinimumDifference(root *TreeNode) int {
	ans := make([]int, 0)
	st := make([]*TreeNode, 0)

	cur := root
	for len(st) > 0 || cur != nil {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
		} else {
			cur = st[len(st)-1]
			st = st[:len(st)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	i, j := 0, 1
	var min int = math.MaxInt64
	for j < len(ans) {
		tmp := ans[j] - ans[i]
		if tmp < min {
			min = tmp
		}
		i++
		j++
	}
	return min
}
