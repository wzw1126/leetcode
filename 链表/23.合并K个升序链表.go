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

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKList(lists, 0, len(lists))
}
func mergeKList(lists []*ListNode, i int, j int) *ListNode {
	m := j - i
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[i]
	}
	left := mergeKList(lists, i, i+m/2)
	right := mergeKList(lists, i+m/2, j)
	return merge(left, right)
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{-1, nil}
	pre := res
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
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
