class Solution {
    public int majorityElement(int[] nums) {
        
        int ans=0,k=0;
        for(int x:nums){
            if(k==0){
                ans=x;
                k=1;
            }else{
                if(x==ans){
                    k++;
                }else{
                    k--;
                }
            }
        }
        return ans;
    }
}