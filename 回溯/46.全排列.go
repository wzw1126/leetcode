func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func([]int)

	dfs = func(nums []int) {

		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true

				dfs(nums)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(nums)
	return res
}

