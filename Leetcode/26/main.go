package main

func removeDuplicates(nums []int) int {
	// k表示不重复数组的下标
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
	return k + 1
}
