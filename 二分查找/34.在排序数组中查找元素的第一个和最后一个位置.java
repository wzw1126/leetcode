class Solution {

    public int[] searchRange(int[] nums, int target) {
        int left=find(nums,target);
        if(left== nums.length || nums[left]!=target) return new int[]{-1,-1};
        int right =find(nums,target+1)-1;
        return new int[]{left,right};
    }

    int find(int[] nums,int target){
        int left=-1,right=nums.length;
        while(left+1<right){
            int mid = (right+left)/2;
            if(nums[mid]<target){
                left=mid;
            }else{
                right=mid;
            }
        }
        return right;
    }
}