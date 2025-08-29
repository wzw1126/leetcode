class Solution {

    List<Integer> path;
    List<List<Integer>> res;
    boolean[] used;

    public List<List<Integer>> permute(int[] nums) {
        path = new ArrayList<>();
        res = new ArrayList<>();
        used = new boolean[nums.length];
        dfs(nums);
        return res;
    }

    public void dfs(int[] nums){
        // 递归终止条件，当路径长度等于数组长度时，将路径加入结果集
        if(path.size() == nums.length){
            res.add(new ArrayList<>(path));
            return;
        }
        for(int i=0;i<nums.length;i++){
            // 剪枝，如果当前数字已经使用过，则跳过
            if(!used[i]){
                path.add(nums[i]);
                used[i]=true;
                dfs(nums);
                    // 回溯，撤销选择
                path.removeLast();
                used[i]=false;
            }
        }
    }
}