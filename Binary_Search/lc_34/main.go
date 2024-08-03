package main

// 找左端点使用模板一，找右端点使用模版二
// 模版一：check(mid) >= x
// 模版二：check(mid) <= x
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	l, r := 0, n-1
	// 找左端点
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	start := l
	if nums[l] != target {
		return []int{-1, -1}
	}

	l, r = 0, n-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	end := l
	return []int{start, end}
}
