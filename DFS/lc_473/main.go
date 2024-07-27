package main

import "sort"

// 思路：
// 1. 从大到小排序，先处理较长火柴的，
//	-> 递归深度低，先处理小的，会导致递归比较深；小的火柴可以拼凑成大的，一开始选择几个小的，拼成功后，留下大的难以选择

// 2. 每条边内部，要求火柴编号递增
// -> 假设数组为[1, 2, 3] 拼接成一条边，那数组[1, 3, 2]也能拼成一条边，这样就重复了，其实应该是同一个答案

// 3. 如果在处理某条边时，某根火柴摆放失败了
// -> 3.1 则跳过长度相同的火柴，因为下一个长度相同的火柴也必然摆放失败
// -> 3.2 如果这根火柴是摆放的第一根，则直接剪掉当前分支。
//        原因：假设这个火柴是a1，这条边拼凑成功，并没有使用a1火柴，但是下一条边也会遇到火柴a1，因为都是按编号查的，那么总会有一条边处理失败，因此拼凑肯定是失败的
// -> 3.3 如果这根火柴是摆放的最后一根，则直接剪掉当前分支
//        原因：cur + nums[i] = target 匹配失败，假设没有使用nums[i]火柴，使用了其他的火柴匹配成功了，那么可以得出
//             cur + nums[u] + ... + nums[k] = target，又因为 cur + nums[i] = target，那么 nums[i] == nums[u] + ... + nums[k]
//             那么nums[u] + ... + nums[k]这些火柴可以被nums[i]一根火柴替换，即 cur + nums[i] = target，可以匹配成功，
//             这个与前提矛盾了，因此使用nums[i]失败后，就直接剪掉当前分支

var st []bool

func makesquare(nums []int) bool {
	n := len(nums)
	st = make([]bool, n)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%4 != 0 {
		return false
	}
	return dfs(nums, 0, 0, sum/4, 0)
}

// start：火柴编号，从start开始遍历
// cur：当前边，拼接长度
// length：每条边拼凑成功的长度
// cnt：第几条边
func dfs(nums []int, start, cur, length, cnt int) bool {
	if cnt == 3 {
		// 如果是第3条边处理成功，则成功了，因为最后一条边的长度已经确定的，总共分成4份，前三份确认了，最后一份就随之确认了
		return true
	}
	if cur == length {
		// 当前边拼凑成功，处理下一条边
		return dfs(nums, 0, 0, length, cnt+1)
	}
	for i := start; i < len(nums); i++ {
		if st[i] {
			continue
		}
		if cur+nums[i] <= length {
			// 这根火柴可以使用
			st[i] = true
			// 继续下一根火柴，这里不能直接 return dfs(nums, i+1, cur + nums[i], length, cnt)
			if dfs(nums, i+1, cur+nums[i], length, cnt) {
				return true
			}
			// 回溯，直接返回就无法回溯了
			st[i] = false
		}
		// nums[i] 火柴拼凑失败，跳过长度相同的
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
		if cur == 0 || cur+nums[i] == length {
			return false
		}
	}
	return false
}
