package main

var used [][]bool

var rows int
var cols int

func exist(board [][]byte, word string) bool {
	rows = len(board)
	cols = len(board[0])
	used = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		used[i] = make([]bool, cols)
	}

	// 这道题不是从(0, 0)出发的
	// return dfs(board, 0, 0, word, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				if dfs(board, i, j, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[index] || used[i][j] {
		return false
	}
	used[i][j] = true
	res := dfs(board, i-1, j, word, index+1) || dfs(board, i, j+1, word, index+1) || dfs(board, i+1, j, word, index+1) || dfs(board, i, j-1, word, index+1)

	used[i][j] = false
	return res
}
