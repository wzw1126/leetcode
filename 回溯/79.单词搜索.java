class Solution {
    boolean[][] visited;
    int m,n;
    public boolean exist(char[][] board, String word) {
        m=board.length;
        n=board[0].length;
        visited=new boolean[m][n];
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(board[i][j]==word.charAt(0) && dfs(board,word,i,j,0)){
                    return true;
                }
            }
        }
        return false;
    }

    boolean dfs(char[][]board,String word,int i,int j,int idx){
        if(board[i][j]!=word.charAt(idx)){
            return false;
        }
        if(idx == word.length()-1){
            return true;
        }
        visited[i][j]=true;
        if(i+1<m && !visited[i+1][j] && dfs(board,word,i+1,j,idx+1)){
            return true;
        }
        if(i-1>=0 && !visited[i-1][j] && dfs(board,word,i-1,j,idx+1)){
            return true;
        }
        if(j+1<n && !visited[i][j+1] && dfs(board,word,i,j+1,idx+1)){
            return true;
        }
        if(j-1>=0 && !visited[i][j-1] && dfs(board,word,i,j-1,idx+1)){
            return true;
        }
        visited[i][j]=false;
        return false;
    }
}