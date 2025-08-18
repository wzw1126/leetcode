// ...existing code...
package 双指针

// trap1 使用双指针（左右指针）方法求接雨水量。
// 思路：维护左右边界的最大高度 leftMax、rightMax。
// 将较矮一侧的指针向内移动，累计该侧可接的雨水（leftMax - height[l] 或 rightMax - height[r]）。
// 时间复杂度：O(n)，空间复杂度：O(1)。
func trap1(height []int) int {
	n := len(height)
	ans := 0
	l, r := 0, n-1
	leftMax, rightMax := 0, 0
	for l < r {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		// 总是移动较矮的一侧，因为短板决定蓄水高度
		if leftMax < rightMax {
			ans += leftMax - height[l]
			l++
		} else {
			ans += rightMax - height[r]
			r--
		}
	}
	return ans
}

// trap 使用前缀最大值与后缀最大值数组计算每个柱子能接的雨水。
// preMax[i] 表示 i 左侧到 i 的最高高度，sufMax[i] 表示 i 到右侧的最高高度。
// 第 i 个位置的蓄水量为 min(preMax[i], sufMax[i]) - height[i]。
// 时间复杂度：O(n)，空间复杂度：O(n)（用于 preMax 和 sufMax）。
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	preMax := make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax := make([]int, n)
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += min(sufMax[i], preMax[i]) - height[i]
	}
	return ans
}
