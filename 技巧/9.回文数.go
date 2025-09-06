//go当中就不能使用x+""了，只能使用strconv.Itoa(x)
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	res := 0
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	//如果是偶数位数，x和res相等
	//如果是奇数位数，res会多一位，把res/10和x比较  12321 运算的结果x=12 res=123  123/10=12
	return x == res || x == res/10
}