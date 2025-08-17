package hot100

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		key := sortKey26(v)        // 自己实现：计数“排序”得到规范 key
		m[key] = append(m[key], v) // 自己实现：map[key] append
	}

	// 自己实现 maps.Values + slices.Collect
	res := make([][]string, 0, len(m))
	for _, group := range m {
		res = append(res, group)
	}
	return res

}
func sortKey26(s string) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	buf := make([]byte, len(s))
	for i := 0; i < 26; i++ {
		for cnt[i] > 0 {
			buf = append(buf, byte('a'+i))
			cnt[i]--
		}
	}
	return string(buf)
}
