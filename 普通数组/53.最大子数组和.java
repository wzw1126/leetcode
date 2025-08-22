/*
*  公式：f[i]=max(f[i-1],0)+nums[i]
*/
class Solution {
    public int maxSubArray(int[] nums) {
        int[] f=new int[nums.length];
        f[0]=nums[0];
        int ans=f[0];
        for(int i=1;i<nums.length;i++){
            f[i]=Math.max(f[i-1],0)+nums[i];
            ans=Math.max(ans,f[i]);
        }
        return ans;
    }
}

//前缀和+贪心
class Solution {
    public int maxSubArray(int[] nums) {
        int ans=Integer.MIN_VALUE;
        int preSum = 0;
        int preMinSum=0;
        for(int x:nums){
            preSum+=x;
            ans=Math.max(ans,preSum-preMinSum);
            preMinSum=Math.min(preSum,preMinSum);
        }
        return ans;
    }
}