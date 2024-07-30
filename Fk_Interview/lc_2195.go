package main

import (
	"math"
	"sort"
)

// 贪心 + 排序
func minimalKSum(nums []int, k int) int64 {
	nums = append(nums, 0, math.MaxInt32)
	sort.Ints(nums)
	sum := 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1] - 1
		if tmp <= 0 {
			continue
		}
		if tmp >= k {
			// for j := 1; j <= k; j++ {
			// 	sum += nums[i-1] + j
			// }
			sum += (nums[i-1] + 1 + nums[i-1] + k) * k / 2
			return int64(sum)
		}
		// [1, 4]，等差数列
		sum += (nums[i-1] + nums[i]) * tmp / 2
		k -= tmp
	}
	return int64(sum)
}

// 错误代码，思路可以看看
func minimalKSum_err(nums []int, k int) int64 {
	arr := make([]int, len(nums)+1)
	arr[0] = 0
	for i := 1; i <= len(nums); i++ {
		arr[i] = nums[i-1]
	}

	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})
	sum := 0
	for i := 1; i < len(arr); i++ {
		tmp := arr[i] - arr[i-1] - 1
		if tmp <= 0 && k == 0 {
			continue
		}
		if tmp > k {
			for j := 1; j <= k; j++ {
				sum += arr[i-1] + j
			}
		} else {
			j := 1
			if j <= tmp {
				for j = 1; j <= tmp; j++ {
					sum += arr[i-1] + j
				}
				k -= tmp
			}
		}
	}
	if k > 0 {
		for j := 1; j <= k; j++ {
			sum += arr[len(arr)-1] + j
		}
	}
	return int64(sum)
}
