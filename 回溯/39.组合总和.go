func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(sum, start int) {
		if sum == target {
			res = append(res, slices.Clone(path))
			return
		}
		if sum > target || start == len(candidates) {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(sum, i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	dfs(0, 0)
	return res
}