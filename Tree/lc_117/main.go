package main

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	var pre *Node
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if i == 0 {
				pre = cur
			}
			if i > 0 {
				pre.Next = cur
				pre = cur
			}
		}
		pre.Next = nil
		queue = queue[n:]
	}
	return root
}
