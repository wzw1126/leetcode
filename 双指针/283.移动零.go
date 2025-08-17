package 双指针

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			tmp := nums[r]
			nums[r] = nums[l]
			nums[l] = tmp
			l++
		}
		r++
	}
}
