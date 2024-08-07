package main

import "fmt"
import "bufio"
import "os"

// 给定正数n，统计n这个数字，x出现的次数
// 1 ～ n 所有数字中，x出现的次数
func count(n, x int) int {
	if n == 0 {
		return 0
	}
	var nums []int
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	// 假设u为1234，那么nums数组为[4,3,2,1]
	// 当x等于0时，x ^ 1 = 1，i从n-2开始，因为最高位是不可能为0的，因此没必要处理
	// 当x不等于0时，x ^ 1 = 0，i从n-1开始
	// 这里不能异或，2 ^ 1 = 3，不是1
	// 改成 1- x，小技巧，这个小技巧也是错的！！！

	n, res, isZero := len(nums), 0, 0
	if x == 0 {
		isZero = 1
	}

	for i := n - 1 - isZero; i >= 0; i-- {

		// 000 ~ abc-1
		// 假设abcdefg，i下标的值对应于d，在nums数组中的顺序是gfedcba，所以要想求abc的值，就是 i+1, n-1
		res += get(nums, n-1, i+1) * power10(i)
		if x == 0 {
			// 必须从001 ~ abc-1
			res -= power10(i)
		}

		if nums[i] == x {
			res += get(nums, i-1, 0) + 1
		} else if nums[i] > x {
			res += power10(i)
		}
	}

	return res
}

// [r, l]区间构成的整数
func get(nums []int, l, r int) int {
	res := 0
	for i := l; i >= r; i-- {
		res = res*10 + nums[i]
	}
	return res
}

func power10(x int) int {
	res := 1
	for x > 0 {
		res = res * 10
		x--
	}
	return res
}

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		var a, b int
		fmt.Fscan(r, &a, &b)

		if a == 0 && b == 0 {
			break
		}

		// 保证[a, b]区间
		if a > b {
			a, b = b, a
		}

		for i := 0; i <= 9; i++ {
			fmt.Printf("%d ", count(b, i)-count(a-1, i))
		}
		fmt.Println()
	}
}
