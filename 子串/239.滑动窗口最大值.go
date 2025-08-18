package 子串

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q := []int{}

	for i, x := range nums {
		//1.右边入
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		//2.左边出
		left := i - k + 1
		if q[0] < left {
			q = q[1:len(q)]
		}

		// 3. 在窗口左端点处记录答案
		if left >= 0 {
			ans[left] = nums[q[0]]
		}
	}
	return ans
}
