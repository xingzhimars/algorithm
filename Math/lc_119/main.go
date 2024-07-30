package main

// 画图

// 1
// 1 1
// 1 2 1
// 1 3 3 1
// 1 4 6 4 1

const N = 34

func getRow(rowIndex int) []int {
	// a := make([][]int, N)
	// for i := 0; i < N; i++ {
	//     a[i] = make([]int, N)
	// }
	// for i := 0; i < N; i++ {
	//     a[i][0] = 1
	// }
	// for i := 1; i < N; i++ {
	//     for j := 1; j < N; j++ {
	//         a[i][j] = a[i-1][j-1] + a[i-1][j]
	//     }
	// }

	// return a[rowIndex][:rowIndex+1]

	// 下面的写法，就是dp二位数组转一位数组，思路一模一样
	a := make([]int, rowIndex+1)
	a[0] = 1
	for i := 1; i <= rowIndex; i++ {
		// 逆序，因为需要上一轮的a[j-1]，那么就必须逆序
		for j := i; j > 0; j-- {
			a[j] += a[j-1]
		}
	}
	return a
}
