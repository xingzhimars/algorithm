package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var ans [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans = make([][]int, 0)
	dfs(root, targetSum, make([]int, 0))
	return ans
}

func dfs(root *TreeNode, targetSum int, path []int) {
	if root == nil {
		return
	}
	// 先加上，容易忽略叶子节点
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		return
	}
	targetSum -= root.Val
	dfs(root.Left, targetSum, path)
	dfs(root.Right, targetSum, path)
	path = path[:len(path)-1]
	targetSum += root.Val
}
