func reverseBits(n int) int {
	res := 0
	for i := 0; i < 32; i++ {
		res = res<<1 | n&1
		n = n >> 1
	}
	return res
}