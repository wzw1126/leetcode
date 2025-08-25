/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums))
}

func dfs(num []int, left int, right int) *TreeNode {
	if left == right {
		return nil
	}
	m := (left + right) / 2
	return &TreeNode{num[m], dfs(num, left, m), dfs(num, m+1, right)}
}
