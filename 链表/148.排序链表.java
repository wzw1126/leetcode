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
    public ListNode sortList(ListNode head) {
        if(head == null || head.next==null){
            return head;
        }
        ListNode head2 = findMid(head);
        head=sortList(head);
        head2=sortList(head2);
        return merge(head,head2);
    }

    public ListNode findMid(ListNode head){
        ListNode pre = head;
        ListNode slow = head;
        ListNode fast = head;
        while(fast!=null && fast.next!=null){
            pre=slow;
            slow=slow.next;
            fast=fast.next.next;
        }
        pre.next=null;
        return slow;
    }

    public ListNode merge(ListNode list1,ListNode list2){
        ListNode dummy = new ListNode(-1,null);
        ListNode pre =dummy;
        while(list1!=null && list2!=null){
            ListNode tmp;
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
        return dummy.next;
    }
}