/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	leftSize := findLeft(inorder, preorder[0])
	pre1 := preorder[1 : 1+leftSize]
	pre2 := preorder[1+leftSize:]
	ino1 := inorder[0:leftSize]
	ino2 := inorder[leftSize+1:]
	l := buildTree(pre1, ino1)
	r := buildTree(pre2, ino2)
	return &TreeNode{preorder[0], l, r}
}

func findLeft(inorder []int, x int) int {
	for i, v := range inorder {
		if v == x {
			return i
		}
	}
	return -1
}