package 滑动窗口

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	cnt := make([]int, 128)
	l, res := 0, 0
	for r := 0; r < n; r++ {
		for cnt[s[r]] > 0 {
			cnt[s[l]]--
			l++
		}
		cnt[s[r]] = 1
		res = max(res, r-l+1)
	}
	return res
}
