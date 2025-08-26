class Solution {
    public int search(int[] nums, int target) {
        int n = nums.length;
        int left =-1 ,right=n;
        while( left+1 < right){
            int mid = (left+right)/2;
            if(nums[mid] == target){
                return mid;
            }
            if(nums[mid]<=nums[n-1]){
                if(nums[mid]<=target && target<=nums[n-1]){
                    left=mid;
                }else{
                    right=mid;
                }
            }else{
                if(nums[left+1]<=target && target<=nums[mid]){
                    right=mid;
                }else{
                    left=mid;
                }
            }
        }
        return -1;
    }
}