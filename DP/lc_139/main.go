package main

// dp[i]：表示以i结尾的子串，是否可以由wordDict划分
func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]struct{})
	for _, v := range wordDict {
		mp[v] = struct{}{}
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			word := s[j:i]
			_, ok := mp[word]
			if ok && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[n]
}
