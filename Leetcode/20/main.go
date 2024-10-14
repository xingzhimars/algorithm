package main

const N = 1e4 + 10

func isValid(s string) bool {
	mp := make(map[byte]byte)
	mp[')'] = '('
	mp[']'] = '['
	mp['}'] = '{'

	st := make([]byte, N)
	tt := 0

	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			tt++
			st[tt] = s[i]
		} else {
			b := mp[s[i]]
			if st[tt] == b {
				tt--
			} else {
				return false
			}
		}
	}
	return tt <= 0
}

func isLeft(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}
