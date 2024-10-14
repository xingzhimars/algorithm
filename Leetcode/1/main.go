package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if k, ok := mp[target-nums[i]]; ok {
			return []int{k, i}
		} else {
			mp[nums[i]] = i
		}
	}
	return []int{}
}
