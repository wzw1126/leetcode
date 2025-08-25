/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) > 0 {
		n := len(que)
		arr := []int{}
		for i := 0; i < n; i++ {
			tmp := que[i]
			if tmp.Left != nil {
				que = append(que, tmp.Left)
			}
			if tmp.Right != nil {
				que = append(que, tmp.Right)
			}
			arr = append(arr, tmp.Val)
		}
		que = que[n:len(que)]
		res = append(res, arr)
	}
	return res
}
