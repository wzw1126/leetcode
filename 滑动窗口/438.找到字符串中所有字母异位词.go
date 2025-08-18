package 滑动窗口

func findAnagrams(s string, p string) []int {
	cnt := [26]int{}
	ans := []int{}
	for _, c := range p {
		cnt[c-'a']++
	}
	left := 0
	for right, c := range s {
		c -= 'a'
		cnt[c]--
		for cnt[c] < 0 {
			cnt[s[left]-'a']++
			left++
		}
		if right-left+1 == len(p) {
			ans = append(ans, left)
		}
	}
	return ans
}
