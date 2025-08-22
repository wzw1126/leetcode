class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int m = matrix.length,n=matrix[0].length;
        int t=0,l=0,r=n-1,d=m-1;
        List<Integer> list = new ArrayList<>();
        while(true){
            for(int i=l;i<=r;i++){
                list.add(matrix[t][i]);
            }
            if(++t>d) break;
            for(int i=t;i<=d;i++){
                list.add(matrix[i][r]);
            }
            if(--r<l)break;
            for(int i=r;i>=l;i--){
                list.add(matrix[d][i]);
            }
            if(--d<t)break;
            for(int i=d;i>=t;i--){
                list.add(matrix[i][l]);
            }
            if(++l>r)break;
        }
        return list;
    }
}