/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode la = headA;
        ListNode lb = headB;
        while(la != lb){
            if(la==null){
                la=headB;
            }else{
                la=la.next;
            }
            if(lb==null){
                lb=headA;
            }else{
                lb=lb.next;
            }
        }
        return la;
    }
}