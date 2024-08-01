package main

// % 3
func singleNumber(nums []int) int {
	a := make([]int, 64)

	for _, v := range nums {
		for i := 0; i < 64; i++ {
			if (v>>i)&1 == 1 {
				// 二进制低位在a[0]
				a[i]++
			}
		}
	}
	ans := 0
	for i := 0; i < 64; i++ {
		if (a[i]%3)&1 == 1 {
			ans += 1 << i
		}
	}
	return ans
}
