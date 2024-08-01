package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 方案一：前缀和 + 哈希表
func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	cnt := make(map[int]int)
	// 前缀和必须有一个s[0]，哨兵
	cnt[0] = 1
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		// 当前节点到根节点的前缀和
		cur += node.Val
		// [i, j]，s[j] - s[i-1] = target，这个区间的节点和为target，那么s[i-1] = s[j] - targte
		// 如果能够找到s[i-1]，不就是说明存在一条从上到下的路径，使得节点和为targetSum吗
		ans += cnt[cur-targetSum]
		// 从根节点到当前节点的前缀和，出现的次数
		cnt[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		cnt[cur]--
	}
	dfs(root, 0)
	return ans
}

// 方案二：dfs
// 直接定义外部变量，var sum int，会存在问题，leetcode的测试用例，很容易污染外部变量
// 这里使用返回值处理
func pathSum_2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := dfs_2(root, targetSum)
	// 以每个节点作为根，进行搜索
	res += pathSum_2(root.Left, targetSum)
	res += pathSum_2(root.Right, targetSum)
	return res
}

// 假设以root为根的子树，进行搜索
func dfs_2(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}

	if targetSum == root.Val {
		res++
	}
	// 由于节点值可能存在负数，因此继续查找
	res += dfs_2(root.Left, targetSum-root.Val)
	res += dfs_2(root.Right, targetSum-root.Val)
	return
}
