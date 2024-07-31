package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 二叉树前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	st := make([]*TreeNode, 0)
	st = append(st, root)

	for len(st) > 0 {
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, cur.Val)

		if cur.Right != nil {
			st = append(st, cur.Right)
		}
		if cur.Left != nil {
			st = append(st, cur.Left)
		}
	}
	return ans
}

// 二叉树中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	st := make([]*TreeNode, 0)
	cur := root

	ans := make([]int, 0)

	for cur != nil || len(st) > 0 {
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
	return ans
}

// 二叉树后序遍历
// 两个栈，s1, s2，s1处理节点，从s1中弹出的节点，压入s2中
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	s1 := make([]*TreeNode, 0)
	s2 := make([]*TreeNode, 0)
	ans := make([]int, 0)

	s1 = append(s1, root)

	for len(s1) > 0 {
		cur := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		// 每次从s1中弹出的节点，都压入s2中
		s2 = append(s2, cur)

		if cur.Left != nil {
			s1 = append(s1, cur.Left)
		}
		if cur.Right != nil {
			s1 = append(s1, cur.Right)
		}
	}

	for len(s2) > 0 {
		cur := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		ans = append(ans, cur.Val)
	}
	return ans
}

// 二叉树的层次遍历，按每层返回
// 还有一个方案，计算每层的个数
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	last := root
	var nextLast *TreeNode

	tmp := make([]int, 0)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		tmp = append(tmp, cur.Val)

		if cur.Left != nil {
			nextLast = cur.Left
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			nextLast = cur.Right
			queue = append(queue, cur.Right)
		}
		if cur == last {
			ans = append(ans, tmp)
			tmp = make([]int, 0)

			last = nextLast
		}
	}
	return ans
}
