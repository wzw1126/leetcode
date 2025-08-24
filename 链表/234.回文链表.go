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

func isPalindrome(head *ListNode) bool {
	midNode := findMid(head)
	head2 := reverse(midNode)
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func reverse(head *ListNode) *ListNode {
	pre := &ListNode{-1, nil}
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = tmp
	}
	return pre.Next
}
