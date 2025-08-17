package 双指针

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		h := 0
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		s := (r - l + 1) * h
		res = max(res, s)
	}
	return res
}
