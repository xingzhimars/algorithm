package main

// x + y是奇数，Alice必赢
// 这种写法，超时
func flowerGame(n int, m int) int64 {
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if (i+j)%2 != 0 {
				ans++
			}
		}
	}
	return int64(ans)
}

// 根据规律发现的
func flowerGame2(n int, m int) int64 {
	res1 := 0
	for j := 1; j <= m; j += 2 {
		res1++
	}
	res2 := 0
	for j := 2; j <= m; j += 2 {
		res2++
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			ans += res1
		} else {
			ans += res2
		}
	}

	return int64(ans)
}

// n内的偶数个数 * m内的奇数个数 + n内的奇数个数 * m内的偶数个数
func flowerGame3(n int, m int) int64 {
	return int64((n/2)*((m+1)/2) + (m/2)*((n+1)/2))
}
