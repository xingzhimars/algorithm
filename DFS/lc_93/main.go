package main

var ans []string

func restoreIpAddresses(s string) []string {
	ans = make([]string, 0)
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, startIndex, pointNum int) {
	if pointNum == 3 {
		if valid(s, startIndex, len(s)-1) {
			ans = append(ans, s)
		}
		return
	}
	for i := startIndex; i < len(s); i++ {
		if valid(s, startIndex, i) {
			s = s[:i+1] + "." + s[i+1:]
			dfs(s, i+2, pointNum+1)
			s = s[:i+1] + s[i+2:]
		}
	}
}

func valid(s string, start, end int) bool {
	if start > end {
		return false
	}
	if start != end && s[start] == '0' {
		return false
	}
	var res int
	for i := start; i <= end; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		res = res*10 + int(s[i]-'0')
	}
	if res > 255 {
		return false
	}
	return true
}
