package main

// 这里的前序遍历一定是#结尾的，
// 当遇到#时，表示子树遍历ok
// 当dfs里面继续调用dfs，说明函数是不该结束的，所以当k == len(s)时，表明递归没有结束，但是数组已经结束了
func isValidSerialization(s string) bool {
	s = s + ","
	k := 0

	var dfs func() bool
	dfs = func() bool {
		// 能递归下去，说明一定有东西，但是此时字符串没有了
		// 如果当#结尾，就不继续递归
		if k == len(s) {
			return false
		}
		// #号结尾，（左或右）子树递归完成
		if s[k] == '#' {
			k += 2
			return true
		}
		for s[k] != ',' {
			k++
		}
		k++
		// 继续判断左子树和右子树
		return dfs() && dfs()
	}
	// 数组用完了，但是递归没有结束
	if !dfs() {
		return false
	}
	// 递归结束了，数组也应该用完，如果数组没有用完，则报错
	return k == len(s)
}
