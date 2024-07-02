package main

var mp map[byte]string
var ans []string

func letterCombinations(digits string) []string {
	ans = make([]string, 0)
	if len(digits) == 0 {
		return ans
	}
	mp = createMap()
	dfs(digits, 0, make([]byte, 0))
	return ans
}

func dfs(digits string, index int, path []byte) {
	if index == len(digits) {
		ans = append(ans, string(path))
		path = make([]byte, 0)
		return
	}
	letter := digits[index]
	letters := mp[letter]
	for i := 0; i < len(letters); i++ {
		path = append(path, letters[i])
		dfs(digits, index+1, path)
		path = path[:len(path)-1]
	}
}

func createMap() map[byte]string {
	mp := make(map[byte]string)
	mp['2'] = "abc"
	mp['3'] = "def"
	mp['4'] = "ghi"
	mp['5'] = "jkl"
	mp['6'] = "mno"
	mp['7'] = "pqrs"
	mp['8'] = "tuv"
	mp['9'] = "wxyz"
	return mp
}
