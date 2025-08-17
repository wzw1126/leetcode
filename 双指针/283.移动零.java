class Solution {
    public void moveZeroes(int[] nums) {
        int l=0,r=0;
        for(;r<nums.length;r++){
            if(nums[r]!=0){
                int tem = nums[r];
                nums[r]=nums[l];
                nums[l]=tem;
                l++;
            }
        }
    }
}


class Solution {
    public void moveZeroes(int[] nums) {
        int l=0,r=0;
        for(;r<nums.length;r++){
            if(nums[r]!=0){
                nums[l]=nums[r];
                l++;
            }
        }
        for(int i=l;i<nums.length;i++){
            nums[i]=0;
        }
    }
}