package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	ans := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]byte, 0)
	}

	row := 0
	flag := -1
	// 将字符串填到对应的位置
	for i := 0; i < n; i++ {
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		ans[row] = append(ans[row], s[i])
		row += flag
	}
	var sb strings.Builder
	for i := 0; i < numRows; i++ {
		sb.Write(ans[i])
	}
	return sb.String()
}

func main() {
	s := "AB"
	res := convert(s, 1)
	fmt.Println(res)
}
