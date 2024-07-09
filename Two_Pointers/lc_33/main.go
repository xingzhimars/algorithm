package main

func search(nums []int, target int) int {
	n := len(nums)

	i, j := 0, n-1
	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == target {
			return mid
		}
		// 左区间
		if nums[mid] > nums[j] {

			// 构造区间[nums[i], target]
			// nums[mid] > target，说明是在区间外的
			// 如果条件改成 if target >= nums[i] && nums[mid] < target，此时有点问题
			if target >= nums[i] && nums[mid] > target {
				j = mid - 1
			} else {
				i = mid + 1
			}

		} else {
			//构造[target, nums[j]]区间，nums[mid]在区间外
			if target <= nums[j] && nums[mid] < target {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}
	return -1
}
