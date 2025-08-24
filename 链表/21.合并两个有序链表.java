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
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        ListNode res = new ListNode(-1,null);
        ListNode pre = res;
        while(list1!=null && list2!=null){
            ListNode tmp = new ListNode();
            if(list1.val<=list2.val){
                tmp=list1.next;
                list1.next=pre.next;
                pre.next=list1;
                list1=tmp;
            }else{
                tmp=list2.next;
                list2.next=pre.next;
                pre.next=list2;
                list2=tmp;
            }
            pre=pre.next;
        }
        if(list1!=null){
            pre.next=list1;
        }
        if(list2!=null){
            pre.next=list2;
        }
        return res.next;
    }
}