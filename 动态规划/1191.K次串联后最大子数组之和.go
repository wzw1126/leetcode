
//如果 k=1，那就是 53 题。注意本题允许子数组为空，答案最小是 0，不可能是负数。
//如果 k=2，我们计算的是 arr+arr 的最大子数组和。直接调用 53 题的代码
//如果 k≥3 呢？
// 定理：设 s 是 arr 的元素和。如果 s>0，那么 arr+arr 的最大子数组和必然横跨两个 arr，不会在其中一个 arr 中间。

func kConcatenationMaxSum(arr []int, k int) int {
	if k == 1 {
		return maxSubArray(arr)
	}
	ans := maxSubArray(append(arr, arr...))
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	ans += max(s, 0) * (k - 2)
	return ans % 1000000007
}
func maxSubArray(nums []int) int {
	fn := 0
	res := fn
	for i := 0; i < len(nums); i++ {
		fn = max(fn, 0) + nums[i]
		res = max(res, fn)
	}
	return res
}