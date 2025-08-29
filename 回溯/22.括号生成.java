class Solution {
    List<String> res;
    List<Character> path;
    public List<String> generateParenthesis(int n) {
        res=new ArrayList<>();
        path = new ArrayList<>();
        dfs(0,0,n);
        return res;
    }

    void dfs(int left,int right,int n){
        if(right==n){
            StringBuilder sb = new StringBuilder();
            for(char ch:path){
                sb.append(ch);
            }
            res.add(sb.toString());
            return;
        }
        if(left<n){
            path.add('(');
            left++;
            dfs(left,right,n);
            left--;
            path.removeLast();
        }
        if(right<left){
            path.add(')');
            right++;
            dfs(left,right,n);
            right--;
            path.removeLast();
        }
    }
}