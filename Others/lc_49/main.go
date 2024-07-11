package main

import "sort"

var ans [][]string

func groupAnagrams(strs []string) [][]string {
	ans = make([][]string, 0)
	mp := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		s := sortStr(strs[i])
		if v, ok := mp[s]; ok {
			v = append(v, strs[i])
			mp[s] = v
		} else {
			mp[s] = []string{strs[i]}
		}
	}

	// 遍历map
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func sortStr(str string) string {
	bs := []byte(str)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	return string(bs)
}
