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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB

	for cur1 != cur2 {
		if cur1 != nil {
			cur1 = cur1.Next
		} else {
			cur1 = headB
		}
		if cur2 != nil {
			cur2 = cur2.Next
		} else {
			cur2 = headA
		}
	}
	return cur1

}
