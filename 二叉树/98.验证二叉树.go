/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//前序遍历
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode)
	res := make([]int, 0)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			res = append(res, root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func isValidBST1(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, low int, upper int) bool {
		if root == nil {
			return true
		}
		return low < root.Val && root.Val < upper && dfs(root.Left, low, root.Val) && dfs(root.Right, root.Val, upper)

	}
	return dfs(root, math.MinInt, math.MaxInt)

}