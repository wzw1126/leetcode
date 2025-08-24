/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		m[cur] = &Node{cur.Val, nil, nil}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		copy := m[cur]
		copy.Next = m[cur.Next]
		copy.Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}
