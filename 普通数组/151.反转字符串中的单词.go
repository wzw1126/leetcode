func reverseWords(s string) string {
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	reverse(b, 0, len(b)-1)
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {

		}
		reverse(b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}