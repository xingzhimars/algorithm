package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = dfs(nums, start, mid-1)
	root.Right = dfs(nums, mid+1, end)
	return root
}
