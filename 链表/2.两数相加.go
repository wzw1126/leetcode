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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{-1, nil}
	la := l1
	lb := l2
	cur := 0
	pre := res
	for la != nil || lb != nil || cur != 0 {
		if la != nil {
			cur += la.Val
			la = la.Next
		}
		if lb != nil {
			cur += lb.Val
			lb = lb.Next
		}
		tmp := &ListNode{cur % 10, nil}
		cur /= 10
		tmp.Next = pre.Next
		pre.Next = tmp
		pre = pre.Next

	}
	return res.Next
}
