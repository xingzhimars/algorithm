package main

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	n := len(a)
	m := len(b)
	i, j := n-1, m-1

	ans := make([]int, n+m)
	idx := n + m - 1

	var carry int
	for i >= 0 && j >= 0 {
		carry += int(a[i] - '0')
		carry += int(b[j] - '0')
		ans[idx] = carry % 2
		carry /= 2
		i--
		j--
		idx--
	}
	for i >= 0 {
		carry += int(a[i] - '0')
		ans[idx] = carry % 2
		carry /= 2

		i--
		idx--
	}

	for j >= 0 {
		carry += int(b[j] - '0')
		ans[idx] = carry % 2
		carry /= 2

		j--
		idx--
	}
	if carry > 0 {
		ans[idx] = carry
	}

	var sb strings.Builder
	for i := idx; i < n+m; i++ {
		tmp := strconv.Itoa(ans[i])
		sb.WriteString(tmp)
	}
	strs := sb.String()
	i = 0
	n = len(strs)
	// 去掉前缀0
	for i < n && strs[i] == '0' {
		i++
	}
	strs = strs[i:]

	// 去掉前缀0后，长度为0的话，就返回"0"
	if len(strs) == 0 {
		return "0"
	}
	return strs
}
