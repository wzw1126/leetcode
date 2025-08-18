// ...existing code...
package 子串

// minWindow 在字符串 s 中找到包含 t 中所有字符的最短子串并返回该子串。
// 思路：滑动窗口 + 哈希表记录每个字符的需求与窗口内计数。
// 时间复杂度：O(|s|+|t|)，空间复杂度：O(字母表大小)
func minWindow(s string, t string) string {
	nt := len(t)
	ns := len(s)
	// 若 t 为空或 s 长度小于 t，则不可能覆盖
	if nt == 0 || ns < nt {
		return ""
	}
	// need: t 中每个字符所需的次数
	need := make(map[byte]int, nt)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// win: 当前窗口中每个字符的次数
	win := make(map[byte]int, nt)
	// valid: 窗口中满足 need 要求的字符种类数
	// left: 窗口左边界索引
	// start/minLen: 记录最短覆盖子串的起始位置与长度
	valid, left, start, minLen := 0, 0, 0, ns+1

	// 右指针不断向右扩展窗口
	for right := 0; right < len(s); right++ {
		v := s[right]
		// 仅当字符在 need 中时才更新窗口计数
		if _, ok := need[v]; ok {
			win[v]++
			// 当某字符在窗口中的计数刚好达到所需次数时，valid 加 1
			if win[v] == need[v] {
				valid++
			}
		}
		// 当所有需要的字符种类都被满足时，尝试收缩左边界以获得最短窗口
		for valid == len(need) {
			// 更新最短窗口信息
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			// 将左边字符移出窗口
			d := s[left]
			left++
			// 仅当移出的字符在 need 中才会影响 valid 和 win
			if _, ok := need[d]; ok {
				// 如果该字符在移出前正好满足需求，移出后 valid 减 1
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	// 未找到符合条件的子串则返回空字符串
	if minLen == ns+1 {
		return ""
	}
	return s[start : start+minLen]
}
