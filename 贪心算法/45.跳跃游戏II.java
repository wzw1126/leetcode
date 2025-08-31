class Solution {
    public int jump(int[] nums) {
        int cur=0,nextCur=0;
        int res=0;
        for(int i=0;i<nums.length-1;i++){
            nextCur=Math.max(nextCur,i+nums[i]);
            if(i==cur){
                cur=nextCur;
                res++;
            }
        }
        return res;
    }
}