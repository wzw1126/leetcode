func maximumCostSubstring(s string, chars string, vals []int) int {
	n := len(s)
	m := make(map[byte]int, len(chars))
	for i := 0; i < len(chars); i++ {
		m[chars[i]-'a'] = vals[i]
	}
	arr := make([]int, n+1)
	maxValue := 0
	for i := 1; i <= n; i++ {
		ch := s[i-1]
		if _, ok := m[ch-'a']; !ok {
			arr[i] = max(arr[i-1], 0) + int(ch-'a'+1)
		} else {
			arr[i] = max(arr[i-1], 0) + m[ch-'a']
		}
		maxValue = max(maxValue, arr[i])
	}
	return maxValue
}