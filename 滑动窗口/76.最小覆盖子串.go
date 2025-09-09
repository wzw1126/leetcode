package main

import "math"

func minWindow(s, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	needKinds := len(need)

	window := make(map[byte]int)
	have := 0

	l := 0
	bestLen := math.MaxInt32
	bestL := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		// 扩窗：只在 need 里的字符才计数
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				have++
			}
		}

		// 满足覆盖条件则尝试缩窗
		for have == needKinds {
			// 更新答案
			if r-l+1 < bestLen {
				bestLen = r - l + 1
				bestL = l
			}

			// 缩窗：移除 s[l]
			d := s[l]
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					have--
				}
				window[d]--
			}
			l++
		}
	}

	if bestLen == math.MaxInt32 {
		return ""
	}
	return s[bestL : bestL+bestLen]
}
