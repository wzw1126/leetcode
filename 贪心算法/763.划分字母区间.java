class Solution {

    public List<Integer> partitionLabels(String S) {
        char[] s=S.toCharArray();
        int n=s.length;
        //先记住每个字母最后出现的位置
        int[] last = new int[26];
        for(int i=0;i<n;i++){
            last[s[i]-'a']=i;
        }
        List<Integer> res = new ArrayList<>();
        //类似跳跃游戏II的思路
        //确定当前区间的右边界,如果到最右边界了，就划分一个区间，并更新左边界
        int start=0,end=0;
        for(int i=0;i<n;i++){
            end=Math.max(end,last[s[i]-'a']);
            if(i==end){
                res.add(end-start+1);
                start=i+1;
            }
        }
        return res;
    }
}