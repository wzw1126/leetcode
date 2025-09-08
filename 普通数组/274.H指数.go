
/*
设 n 为 citations 的长度，即这名研究者发表的论文数。根据题意，h 不可能超过 n，所以对于引用次数大于 n 的论文，我们在统计的时候，可以看成是引用次数等于 n 的论文。
例如 n=5，假设 h 是 5，那么无论引用次数是 6 还是 5，都满足 ≥5，所以 6 可以看成是 5，毕竟我们只需要统计有多少个数字是 ≥5 的。
所以，创建一个长为 n+1 的 cnt 数组，统计 min(citations[i],n) 的出现次数。
设 s 为引用次数 ≥i 的论文数，我们需要算出满足 s≥i 的最大的 i。
*/
func hIndex(citations []int) int {
	n := len(citations)
	arr := make([]int, n+1)
	for _, v := range citations {
		arr[min(v, n)]++
	}
	s := 0
	for i := n; ; i-- {
		s += arr[i]
		if s >= i {
			return i
		}
	}
}