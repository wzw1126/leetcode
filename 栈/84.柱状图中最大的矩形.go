/*
	面积最大矩形的高度一定是 heights 中的元素。
	这可以用反证法证明：假如高度不在 heights 中，比如 4，那我们可以增加高度直到触及某根柱子的顶部，比如增加到 5，
	由于矩形底边长不变，高度增加，我们得到了面积更大的矩形，矛盾，所以面积最大矩形的高度一定是 heights 中的元素。
	枚举每个 h=heights[i]，作为矩形的高。那么矩形的宽最大是多少？我们需要知道：
	在 i 左侧的小于 h 的最近元素的下标 left，如果不存在则为 −1。求出了 left，那么 left+1 就是矩形最左边那根柱子。
	在 i 右侧的小于 h 的最近元素的下标 right，如果不存在则为 n。求出了 right，那么 right−1 就是矩形最右边那根柱子。

*/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	st := []int{}
	for i, h := range heights {
		for len(st) > 0 && heights[st[len(st)-1]] >= h {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			left[i] = -1
		} else {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	right := make([]int, n)
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && heights[st[len(st)-1]] >= heights[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			right[i] = n
		} else {
			right[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	ans := 0
	for i, h := range heights {
		ans = max(ans, h*(right[i]-left[i]-1))
	}
	return ans
}