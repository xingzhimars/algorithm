package main

var ans []string

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	dfs(n, n, "")
	return ans
}

func dfs(left, right int, path string) {
	if left == 0 && right == 0 {
		ans = append(ans, path)
		return
	}
	if left > right {
		return
	}
	if left > 0 {
		dfs(left-1, right, path+"(")
	}
	if right > 0 {
		dfs(left, right-1, path+")")
	}
}
