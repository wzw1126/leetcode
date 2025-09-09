func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	res := 0
	for i := 0; i < n; {
		if i+1 < n && m[s[i]] < m[s[i+1]] {
			res += m[s[i+1]] - m[s[i]]
			i += 2
		} else {
			res += m[s[i]]
			i++
		}
	}
	return res
}