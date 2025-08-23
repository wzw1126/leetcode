/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package 链表
type ListNode struct {
   Val int
   Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{Val: -1, Next: head}
	cur := head
	pre.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = tmp
	}
	return pre.Next
}
