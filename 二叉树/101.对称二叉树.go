/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isFlag(root.Left, root.Right)
}

func isFlag(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && isFlag(left.Left, right.Right) && isFlag(left.Right, right.Left)
}