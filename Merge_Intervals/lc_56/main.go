package main

import "sort"

func merge(intervals [][]int) [][]int {
	// 按照左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		pre := res[len(res)-1]

		if pre[1] >= cur[0] {
			pre[1] = max(pre[1], cur[1])
		} else {
			res = append(res, cur)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
