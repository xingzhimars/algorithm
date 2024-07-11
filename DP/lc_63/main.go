package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])

	// 定义成 全局var dp [N][N]int，就出错
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	if obstacleGrid[rows-1][cols-1] == 1 {
		return 0
	}

	for i := 0; i < rows; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 0; i < cols; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			var v int
			if obstacleGrid[i-1][j] == 0 {
				v += dp[i-1][j]
			}
			if obstacleGrid[i][j-1] == 0 {
				v += dp[i][j-1]
			}
			dp[i][j] = v
		}
	}
	return dp[rows-1][cols-1]
}
