//思想：贪心+数学
//从左到右扫描ratings数组，找到每一个山峰，然后计算出这个山峰对应的糖果数
func candy(ratings []int) int {
	n := len(ratings)
	ans := n
	for i := 0; i < n; i++ {
		start := i
		//这个if的作用是处理平峰的情况，比如1 0 2这种情况
		//第一次计算可以得到start = 0,top=0得到inc=0,dec=1,得到ans=1
		//第二次计算可以得到start = 2,top=2得到inc=0,dec=0,得到ans=1,但实际上ratings[2]>ratings[1]，所以start应该减1，是一个递增区间
		if i > 0 && ratings[i-1] < ratings[i] {
			start--
		}
		for i+1 < n && ratings[i] < ratings[i+1] {
			i++
		}
		top := i
		for i+1 < n && ratings[i] > ratings[i+1] {
			i++
		}
		inc := top - start
		dec := i - top
		ans += (inc*(inc-1)+dec*(dec-1))/2 + max(inc, dec)
	}
	return ans
}