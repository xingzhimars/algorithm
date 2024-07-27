package main

func removeElement(nums []int, val int) int {
	n := len(nums)
	j := n - 1
	i := 0
	for i <= j {
		if nums[i] != val {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return i
}
