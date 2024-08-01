package main

// 方案一：
func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0]++

	preSum := make([]int, len(nums)+1)

	preSum[0] = 0
	ans := 0

	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	for i := 1; i <= len(nums); i++ {
		ans += mp[preSum[i]-k]
		mp[preSum[i]]++
	}
	return ans
}

// 方案二：
func subarraySum_2(nums []int, k int) int {
	// 前缀和
	n := len(nums)

	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	count := 0
	// [i, j]
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if preSum[j]-preSum[i-1] == k {
				count++
			}
		}
	}
	return count
}
