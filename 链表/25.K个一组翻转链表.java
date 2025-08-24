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
    public ListNode reverseKGroup(ListNode head, int k) {
        if(head==null || k<=1) return head;
        ListNode dummy = new ListNode(-1,head);
        ListNode pre = dummy;
        while(true){
            ListNode tail = pre;
            for(int i=0;i<k;i++){
                tail=tail.next;
                if(tail==null){
                    return dummy.next;
                }
            }
            ListNode nextGroupHead = tail.next;

            ListNode prev = nextGroupHead;
            ListNode cur = pre.next;
            while(cur!=nextGroupHead){
                ListNode tmp = cur.next;
                cur.next=prev;
                prev = cur;
                cur = tmp;
            }
            ListNode oldHead = pre.next;
            pre.next=prev;
            pre = oldHead;
        }
    }
}