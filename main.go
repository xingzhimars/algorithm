package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//nums := []int{1, 4, 25, 10, 25}
	nums := []int{5, 6}
	k := 6
	res := minimalKSum(nums, k)
	fmt.Println(res)
}

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
			for j := 1; j <= k; j++ {
				sum += nums[i-1] + j
			}
			return int64(sum)
		}
		sum += (nums[i-1] + nums[i]) * tmp / 2
		k -= tmp
	}
	return int64(sum)
}

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
