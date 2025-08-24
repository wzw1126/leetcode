/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{-1, head}
	pre := res
	for pre.Next != nil && pre.Next.Next != nil {
		tmp := pre.Next.Next.Next
		tmp2 := pre.Next.Next
		tmp2.Next = pre.Next
		pre.Next.Next = tmp
		pre.Next = tmp2
		pre = pre.Next.Next
	}
	return res.Next
}
