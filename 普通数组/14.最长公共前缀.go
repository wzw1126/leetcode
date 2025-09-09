func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for i := 1; i < n; i++ {
		minLen = min(minLen, len(strs[i]))
	}
	for i := minLen; i >= 0; i-- {
		str := strs[0][:i]
		j := 1
		for ; j < n; j++ {
			if str != strs[j][:i] {
				break
			}
		}
		if j == n {
			return str
		}
	}
	return ""
}