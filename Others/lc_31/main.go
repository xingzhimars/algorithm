package main

func nextPermutation(nums []int) {
	n := len(nums)

	j := n - 1
	i := j - 1
	for i >= 0 {
		if nums[i] < nums[j] {
			break
		}
		i--
		j--
	}
	if i >= 0 {
		k := n - 1
		// for k >= j && nums[k] >= nums[j] {
		//     k--
		// }
		// k++

		// 找到比nums[j]大的，就直接返回
		for k >= j {
			if nums[k] > nums[j] {
				break
			}
			k--
		}

		nums[j], nums[k] = nums[k], nums[j]
	}
	p1, p2 := j, n-1
	for p1 <= p2 {
		nums[p1], nums[p2] = nums[p2], nums[p1]
		p1++
		p2--
	}
}
