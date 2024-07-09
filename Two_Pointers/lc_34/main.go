package main

func searchRange(nums []int, target int) []int {
	n := len(nums)
	i, j := 0, n-1

	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			p1 := mid
			for p1 >= i {
				if nums[p1] != nums[mid] {
					break
				}
				p1--
			}

			p2 := mid
			for p2 <= j {
				if nums[p2] != nums[mid] {
					break
				}
				p2++
			}
			return []int{p1 + 1, p2 - 1}
		}
	}
	return []int{-1, -1}
}
