//递归+记录返回值=记忆化搜索
class Solution {
    public int climbStairs(int n) {
        int[] memo=new int[n+1];
        return dfs(memo,n);
    }
    int dfs(int[] memo,int n){
        if(n == 0 || n==1){
            return 1;
        }
        if(memo[n]!=0){
            return memo[n];
        }
        memo[n]=dfs(memo,n-1)+dfs(memo,n-2);
        return memo[n];
    }
}

//递推
class Solution {
    public int climbStairs(int n) {
        int[] f = new int[n + 1];
        f[0] = f[1] = 1;
        for (int i = 2; i <= n; i++) {
            f[i] = f[i - 1] + f[i - 2];
        }
        return f[n];
    }
}
//递推空间优化
class Solution {
    public int climbStairs(int n) {
        int f0 = 1;
        int f1 = 1;
        for (int i = 2; i <= n; i++) {
            int newF = f1 + f0;
            f0 = f1;
            f1 = newF;
        }
        return f1;
    }
}
