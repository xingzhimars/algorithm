package main

func removeDuplicates(nums []int) int {
	return process(nums, 2)
}

// 能否将a[i]放入数组？ 就看a[i-2]是否等于a[i]，如果等于，则不能放，这个前提是有序的
// 定义u表示新数组的下标，如果u位置想要放元素的话，就看a[u-2]是否等于想放的元素nums[i]
func process(nums []int, k int) int {
	u := 0
	for i := 0; i < len(nums); i++ {
		if u < k || nums[u-k] != nums[i] {
			nums[u] = nums[i]
			u++
		}
	}
	return u
}
