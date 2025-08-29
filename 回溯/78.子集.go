func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(int)
	dfs = func(start int) {
		if start == len(nums) {
			tmp := make([]int, len(path))
			for i, v := range path {
				tmp[i] = v
			}
			res = append(res, tmp)
			return
		}
		dfs(start + 1)
		path = append(path, nums[start])
		dfs(start + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return res

}