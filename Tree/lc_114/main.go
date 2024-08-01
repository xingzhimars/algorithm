package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	// for cur != nil {
	//     node := cur.Left

	//     for node != nil && node.Right != nil {
	//         node = node.Right
	//     }
	//     if node != nil {
	//         node.Right = cur.Right
	//     }
	//     cur.Right = cur.Left
	//     cur.Left = nil

	//     cur = cur.Right
	// }

	for cur != nil {
		tmp := cur.Left
		if tmp != nil {
			for tmp != nil && tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
