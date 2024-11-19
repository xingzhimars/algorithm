package main

// dp[i][j]：以nums1[i-1]和nums2[j-1]结尾的数组的公共子数组的长度
// 如果nums1[i-1] == nums2[j-1]，dp[i][j] = dp[i-1][j-1] + 1
func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}

	// 最大长度不一定以nums1[n1-1]结尾，因此这里需要遍历
	res := 0
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			res = max(res, dp[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
