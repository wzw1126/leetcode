class Solution {
    public int majorityElement(int[] nums) {
        int cnt=0,index=0;
        for(int i=0;i<nums.length;i++){ 
            if(cnt==0){
                index=i;
            }
            if(nums[i]!=nums[index]){
                cnt--;
            }else{
                cnt++;
            }
        }
        return nums[index];
    }
}