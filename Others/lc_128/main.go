package main

// 思路：哈希表
// 遍历数组，当前元素是cur，如果cur - 1存在于数组，则跳过，不过就是重复计算
// for循环判断cur+1是否在数组中，计算出最长距离
func longestConsecutive(nums []int) int {
	var maxLen int
	mp := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = struct{}{}
	}
	for i := 0; i < len(nums); i++ {
		//_, ok := mp[nums[i]]
		//if ok {
		//	_, ok2 := mp[nums[i]-1]
		//	if ok2 {
		//		continue
		//	} else {
		//		cnt := 0
		//		v := nums[i]
		//		for {
		//			if _, ok3 := mp[v]; ok3 {
		//				cnt++
		//				v += 1
		//			} else {
		//				break
		//			}
		//		}
		//		maxLen = max(maxLen, cnt)
		//	}
		//}
		cur := nums[i]
		_, ok := mp[cur-1]
		if ok {
			continue
		}
		for {
			if _, ok := mp[cur+1]; ok {
				cur = cur + 1
			} else {
				break
			}
		}
		maxLen = max(maxLen, cur-nums[i]+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
