/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	sum, cnt := 0, 0
	var dfs func(*TreeNode, int, int)
	dfs = func(root *TreeNode, targetSum int, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		if sum == targetSum {
			cnt++
		}
		dfs(root.Left, targetSum, sum)
		dfs(root.Right, targetSum, sum)
	}

	que := []*TreeNode{root}
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			tmp := que[i]
			dfs(tmp, targetSum, sum)
			if tmp.Left != nil {
				que = append(que, tmp.Left)
			}
			if tmp.Right != nil {
				que = append(que, tmp.Right)
			}
		}
		que = que[size:]
	}
	return cnt
}   

