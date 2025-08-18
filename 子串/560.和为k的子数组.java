class Solution {
    public int subarraySum(int[] nums, int k) {
        Map<Integer,Integer>map = new HashMap<>();
        int res=0;
        int sum=0;
        map.put(0,1);
        for(int x:nums){
            sum+=x;
            res+=map.getOrDefault(sum-k,0);
            map.put(sum,map.getOrDefault(sum,0)+1);
        }
        return res;

    }
}