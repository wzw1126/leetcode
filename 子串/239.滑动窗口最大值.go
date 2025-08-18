// ...existing code...
package 子串

// maxSlidingWindow 返回数组 nums 中每个大小为 k 的滑动窗口的最大值。
// 思路：使用单调双端队列（存放元素索引，队列从头到尾对应的值单调递减）
// 1) 新元素入队时，从队尾弹出所有比它小或等于的元素索引，保证队头始终是窗口最大值的索引。
// 2) 当队头索引越过窗口左边界时，将其弹出。
// 时间复杂度：O(n)，每个元素最多进出队列一次；空间复杂度：O(k)。
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q := []int{} // 存放索引的单调队列，队头为当前窗口最大值的索引

	for i, x := range nums {
		// 1. 新元素从右端入队前，弹出所有值 <= x 的索引，保持单调递减
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// 2. 计算当前窗口的左边界，若队头索引小于 left，则说明已移出窗口，从头弹出
		left := i - k + 1
		if len(q) > 0 && q[0] < left {
			q = q[1:len(q)]
		}

		// 3. 当左边界 >= 0 时，窗口已形成，队头即为当前窗口的最大值索引
		if left >= 0 {
			ans[left] = nums[q[0]]
		}
	}
	return ans
}
