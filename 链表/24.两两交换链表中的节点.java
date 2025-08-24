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
    public ListNode swapPairs(ListNode head) {
        ListNode res = new ListNode(-1,head);
        ListNode pre = res;
        while(pre.next!=null && pre.next.next!=null){
            ListNode tmp = pre.next.next.next;
            ListNode tmp2 = pre.next.next;
            pre.next.next.next=pre.next;
            pre.next.next=tmp;
            pre.next=tmp2;
            pre=pre.next.next;
        }
        return res.next;
    }
}