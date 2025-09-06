func addBinary(a string, b string) string {
	cnt, i, j := 0, len(a)-1, len(b)-1
	arr := []byte{}
	for i >= 0 || j >= 0 || cnt > 0 {
		if i >= 0 {
			cnt += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			cnt += int(b[j] - '0')
			j--
		}
		arr = append(arr, byte(cnt%2)+'0')
		cnt /= 2
	}
	swapArray(arr)
	return string(arr)
}

func swapArray(nums []byte) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}