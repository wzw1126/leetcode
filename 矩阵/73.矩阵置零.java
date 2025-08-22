class Solution {
    public void setZeroes(int[][] matrix) {
        int m=matrix.length,n=matrix[0].length;
        Map<Integer,Integer>mapi = new HashMap<>();
        Map<Integer,Integer>mapj = new HashMap<>();
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(matrix[i][j]==0){
                    mapi.put(i,mapi.getOrDefault(i,0)+1);
                    mapj.put(j,mapj.getOrDefault(j,0)+1);
                }
            }
        }
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(mapi.containsKey(i) || mapj.containsKey(j))matrix[i][j]=0;
            }
        }
    }
}