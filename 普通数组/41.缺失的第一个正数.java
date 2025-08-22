class Solution {
    public int firstMissingPositive(int[] nums) {
        Map<Integer,Integer> map = new HashMap<>();
        for(int x:nums){
            map.put(x,map.getOrDefault(x,0)+1);
        }
        int k=1;
        for(;;k++){
            if(!map.containsKey(k)){
                break;
            }
        }
        return k;
    }
}