func maxProfit(prices []int) int {
	res := 0
	preMin := prices[0]
	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-preMin)
		preMin = min(preMin, prices[i])
	}
	return res
}