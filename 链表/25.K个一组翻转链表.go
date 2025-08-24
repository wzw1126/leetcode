// Go 版本
// Definition:
//
//	type ListNode struct {
//	    Val  int
//	    Next *ListNode
//	}
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy

	for {
		// 找到当前组的尾节点 tail，若不足 k 个则返回结果
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nextGroupHead := tail.Next

		// 原地反转 [pre.Next, tail]（使用 prev 指向 nextGroupHead 便于尾部连接）
		prev := nextGroupHead
		cur := pre.Next
		for cur != nextGroupHead {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}

		// 接回反转后的部分，并移动 pre 到已处理段的尾部
		oldHead := pre.Next
		pre.Next = prev
		pre = oldHead
	}
}
