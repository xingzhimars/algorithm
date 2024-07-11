package main

func rotate(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	i, j := 0, rows-1
	for i <= j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i <= j {
				continue
			}
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
