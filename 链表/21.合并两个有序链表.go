/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{-1, nil}
	pre := res
	for list1 != nil && list2 != nil {
		tmp := &ListNode{}
		if list1.Val <= list2.Val {
			tmp = list1.Next
			list1.Next = pre.Next
			pre.Next = list1
			list1 = tmp
		} else {
			tmp = list2.Next
			list2.Next = pre.Next
			pre.Next = list2
			list2 = tmp
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return res.Next
}
