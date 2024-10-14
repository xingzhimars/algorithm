package main

// 思路：在两个有序数组中，找到第K个大小的数
// 若A[k/2] < B[k/2]，那么A[0...k/2]这部分元素肯定不是中位数，后面就在A[k/2+1...len(A)-1]和B数组中继续找
// 如果idx1 + k/2，超出了A的范围，那么中位数就在数组B中，即B[idx2+k-1]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	len := m + n
	var val float64
	if len%2 == 0 {
		num := findK(nums1, 0, nums2, 0, len/2) + findK(nums1, 0, nums2, 0, len/2+1)
		val = float64(num) / float64(2)
	} else {
		num := findK(nums1, 0, nums2, 0, len/2+1)
		val = float64(num)
	}
	return val
}

// k：第几个，不是下标
func findK(nums1 []int, idx1 int, nums2 []int, idx2, k int) int {
	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		half := k / 2
		newIdx1 := min(idx1+half, len(nums1)) - 1
		newIdx2 := min(idx2+half, len(nums2)) - 1
		if nums1[newIdx1] < nums2[newIdx2] {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}
}

func findK2(nums1 []int, idx1 int, nums2 []int, idx2, k int) int {
	if idx1 == len(nums1) {
		return nums2[idx2+k-1]
	}
	if idx2 == len(nums2) {
		return nums1[idx1+k-1]
	}
	if k == 1 {
		return min(nums1[idx1], nums2[idx2])
	}
	half := k / 2
	newIdx1 := min(idx1+half, len(nums1)) - 1
	newIdx2 := min(idx2+half, len(nums2)) - 1

	if nums1[newIdx1] < nums2[newIdx2] {
		k -= newIdx1 - idx1 + 1
		idx1 = newIdx1 + 1
		return findK2(nums1, idx1, nums2, idx2, k)
	} else {
		k -= newIdx2 - idx2 + 1
		idx2 = newIdx2 + 1
		return findK2(nums1, idx1, nums2, idx2, k)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
