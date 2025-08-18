package 子串

func subarraySum(nums []int, k int) int {
	cnt := make(map[int]int, len(nums))
	sum := 0
	ans := 0
	cnt[0] = 1
	for _, v := range nums {
		sum += v
		ans += cnt[sum-k]
		cnt[sum]++
	}
	return ans
}
