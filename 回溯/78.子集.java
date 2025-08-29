class Solution {
    List<List<Integer>> res;
    List<Integer> path;
    public List<List<Integer>> subsets(int[] nums) {
        res = new ArrayList<>();
        path = new ArrayList<>();
        dfs(nums,0);
        return res;
    }

    public void dfs(int[] nums,int start){
        if(start == nums.length){
            res.add(new ArrayList<>(path));
            return ;
        }
        dfs(nums,start+1);
        path.add(nums[start]);
        dfs(nums,start+1);
        path.removeLast();
    }
}