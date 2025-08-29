/*
      1.从右往左找第一个“下降位”索引 i：nums[i] < nums[i+1]（注意返回的是 i 而不是 i+1）。
      2.在右侧尾部找到比 nums[i] 大的最小元素（从右往左第一个 > nums[i] 的位置 j），交换 nums[i] 和 nums[j].
	  3.再把 i+1..end 整段反转（使之最小化）。
    */
class Solution {
    public void nextPermutation(int[] nums) {
        int n = nums.length;
        int i=n-2;
        for(;i>=0;i--){
            if(nums[i+1]>nums[i]){
                break;
            }
        }
        if(i>=0){
            int j=n-1;
            for(;j>=0 && nums[j]<=nums[i];){
                j--;
            }
            int temp = nums[i];
            nums[i]=nums[j];
            nums[j]=temp;
        }
        for(int k=i+1,m=n-1;k<m;k++,m--){
            int temp=nums[k];
            nums[k]=nums[m];
            nums[m]=temp;
        }
    }
}