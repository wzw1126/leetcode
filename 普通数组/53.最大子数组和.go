package 普通数组

import "math"

func maxSubArray1(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1], 0) + nums[i]
		ans = max(ans, f[i])
	}
	return ans
}

func maxSubArray(nums []int) int {
	preSum, preMinSum, ans := 0, 0, math.MinInt
	for _, v := range nums {
		preSum += v
		ans = max(ans, preSum-preMinSum)
		preMinSum = min(preMinSum, preSum)
	}
	return ans
}
