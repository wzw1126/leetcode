package 双指针

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0, n-2)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		for j, k := i+1, n-1; j < k; {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp > 0 {
				k--
			} else if tmp < 0 {
				j++
			} else {
				arr := []int{nums[i], nums[j], nums[k]}
				res = append(res, arr)
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return res
}
