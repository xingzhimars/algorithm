package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
