func majorityElement(nums []int) int {
	num, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 {
			num = v
			cnt++
		} else if v == num {
			cnt++
		} else {
			cnt--
		}
	}
	return num
}