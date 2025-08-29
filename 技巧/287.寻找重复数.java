 /** [1,3,4,2,2]当作链表来
  * 1.快慢指针找到相遇点
  * 2.从相遇点和起点同时出发，找到入环点
  * 3.返回入环点
  * 即 1 3 4 2 2
  *    0 1 2 3 4
  * 链表： 0->1->3->2->4->2这里有个环
  * 
  */
  

class Solution {
    public int findDuplicate(int[] nums) {
        int slow=0;
        int fast=0;
        slow=nums[slow];
        fast=nums[nums[fast]];
        while(slow!=fast){
            slow=nums[slow];
            fast=nums[nums[fast]];
        }
        int pre1=0;
        int pre2=slow;
        while(pre1!=pre2){
            pre1=nums[pre1];
            pre2=nums[pre2];
        }
        return pre1;
    }
}