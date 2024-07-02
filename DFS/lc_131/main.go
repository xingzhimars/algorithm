package main

var ans [][]string

const N = 20

var dp [N][N]bool

// 1. 记录每个子串是否是回文子串（找最长回文子串的题目）
// 2. 遍历所有子串，判断是否是回文子串
// 题目中讲到所有可能的情况，基本都是dfs
func partition(s string) [][]string {
	ans = make([][]string, 0)

	n := len(s)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i >= j {
				dp[i][j] = true
			}
		}
	}

	for k := 1; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	dfs(s, 0, make([]string, 0))
	return ans
}

// 所有可能的方案，dfs
func dfs(s string, index int, path []string) {
	if index == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		return
	}
	for i := index; i < len(s); i++ {
		if !dp[index][i] {
			continue
		}
		path = append(path, s[index:i+1])
		dfs(s, i+1, path)
		path = path[:len(path)-1]
	}
}
