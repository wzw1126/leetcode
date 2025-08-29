class Solution {
    List<List<Integer>> res;
    List<Integer> path;
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        res=new ArrayList<>();
        path=new ArrayList<>();
        dfs(candidates,0,0,target);
        return res;
    }

    void dfs(int[] candidates,int start,int sum,int target){
        if(sum == target){
            res.add(new ArrayList<>(path));
            return;
        }
        if(sum > target || start==candidates.length){
            return;
        }
        for(int i=start;i<candidates.length;i++){
            path.add(candidates[i]);
            sum+=candidates[i];
            dfs(candidates,i,sum,target);
            path.removeLast();
            sum-=candidates[i];
        }
    }
}