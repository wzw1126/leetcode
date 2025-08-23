/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode pre = new ListNode(-1,head);
        ListNode cur = head;
        pre.next=null;
        while(cur!=null){
            ListNode tmp = cur.next;
            cur.next=pre.next;
            pre.next=cur;
            cur=tmp;
        }
        return pre.next;
    }
}