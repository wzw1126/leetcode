class Solution {
    public int longestConsecutive(int[] nums) {
        if(nums == null || nums.length == 0 ) return 0;
        Set<Integer> set = new HashSet<>(nums.length *2);
        for(int x:nums){
            set.add(x);
        }
        int ans=0;
        for(int x:nums){
            if(set.contains(x-1)) continue;
            int cur=x,len=1;
            while(set.contains(cur+1)){
                cur++;
                len++;
            }
            ans = Math.max(ans,len);
            if(ans*2 >= nums.length){
                break;
            }
        }
        return ans;
    }
}