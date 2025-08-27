//思路：使用哈希表统计频次，然后将哈希表的entry放入list中排序，最后取前k个元素
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	type kv struct {
		key int
		val int
	}
	pairs := make([]kv, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val > pairs[j].val
	})
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, pairs[i].key)
	}
	return res
}