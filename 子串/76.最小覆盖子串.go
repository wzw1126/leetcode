package 子串

func minWindow(s string, t string) string {
	nt := len(t)
	ns := len(s)
	if nt == 0 || ns < nt {
		return ""
	}
	need := make(map[byte]int, nt)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	win := make(map[byte]int, nt)
	valid, left, start, minLen := 0, 0, 0, ns+1
	for right := 0; right < len(s); right++ {
		v := s[right]
		if _, ok := need[v]; ok {
			win[v]++
			if win[v] == need[v] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	if minLen == ns+1 {
		return ""
	}
	return s[start : start+minLen]
}
