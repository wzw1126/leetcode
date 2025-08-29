class Solution {
    List<List<String>> res;
    List<String>path;
    public List<List<String>> partition(String s) {
        res = new ArrayList<>();
        path = new ArrayList<>();
        dfs(s,0);
        return res;
    }
    void dfs(String s,int start){
        if(start == s.length()){
            res.add(new ArrayList<>(path));
            return;
        }
        for(int i=start;i<s.length();i++){
            if(isPalind(s,start,i)){
                String str = s.substring(start,i+1);
                path.add(str);
                dfs(s,i+1);
                path.removeLast();
            }
        }
    }

    boolean isPalind(String s,int left,int right){
        while(left<right){
            if(s.charAt(left++)!=s.charAt(right--)){
                return false;
            }
        }
        return true;
    }
}