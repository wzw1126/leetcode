
//闭区间二分查找
class Solution {
    public int searchInsert(int[] nums, int target) {
        int left=0,right=nums.length-1;
        while(left<=right){
            int mid = left+(right-left)/2;
            if(nums[mid]<target){
                left=mid+1;
            }else{
                right=mid-1;
            }
        }
        return left;
    }
}


//开区间二分查找
class Solution {
    public int searchInsert(int[] nums, int target) {
        int left=-1,right=nums.length;
        while(left+1 < right){
            int mid = left+(right-left)/2;
            if(nums[mid]<target){
                left=mid;
            }else{
                right=mid;
            }
        }
        return right;
    }
}