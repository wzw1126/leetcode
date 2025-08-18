package 双指针

func trap1(height []int) int {
	n := len(height)
	ans := 0
	l, r := 0, n-1
	leftMax, rightMax := 0, 0
	for l < r {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
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

func trap(height []int) int {
	n := len(height)
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
