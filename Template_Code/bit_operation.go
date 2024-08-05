package main

// 位运算

// x的二进制表示中第k位是几
func solve(x, k int) int {
	return (x >> k) & 1
}

// 返回x的最后一位
// 使用lowBit，可以计算二进制中1的个数
func lowBit(x int) int {
	return x & (-x)
}
