class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int n = cost.length;
        int[] memo = new int[n+1];
        Arrays.fill(memo,-1);
        return dfs(cost,memo,n);
    }
    public int dfs(int[] cost,int[] memo,int i){
        if(i<2){
            return 0;
        }
        if(memo[i]!=-1){
            return memo[i];
        }
        return memo[i]=Math.min(cost[i-1]+dfs(cost,memo,i-1),cost[i-2]+dfs(cost,memo,i-2));
    }
}