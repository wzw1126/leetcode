//思路：使用哈希表统计频次，然后将哈希表的entry放入list中排序，最后取前k个元素
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer,Integer> map = new HashMap<>();
        for(int x:nums){
            map.put(x,map.getOrDefault(x,0)+1);
        }
        List<Map.Entry<Integer,Integer>> list = new ArrayList<>(map.entrySet());
        list.sort((a,b)->b.getValue() - a.getValue());
        int[] res = new int[k];
        for(int i=0;i<k;i++){
            res[i]=list.get(i).getKey();
        }
        return res;
    }
}