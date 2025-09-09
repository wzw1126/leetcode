func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	n := len(needle)
	for i := 0; i <= len(haystack)-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}