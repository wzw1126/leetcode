func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1
	for i >= 0 {
		if s[i] != ' ' {
			break
		}
		i--
	}
	start := i
	for start >= 0 {
		if s[start] == ' ' {
			break
		}
		start--
	}
	return i - start
}