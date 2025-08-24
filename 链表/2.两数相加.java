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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode la = l1;
        ListNode lb = l2;
        ListNode res = new ListNode(-1,null);
        ListNode pre=res;
        int cur=0;
        while(la!=null || lb!=null || cur!=0){
            if(la!=null){
                cur+=la.val;
                la=la.next;
            }
            if(lb!=null){
                cur+=lb.val;
                lb=lb.next;
            }
            ListNode node = new ListNode(cur%10,null);
            cur/=10;
            node.next=pre.next;
            pre.next=node;
            pre=pre.next;
        }
        return res.next;

    }
}