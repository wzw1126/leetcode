/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public int pathSum(TreeNode root, int targetSum) {
        if(root==null) return 0;
        Queue<TreeNode> que = new LinkedList<>();
        que.add(root);
        while(!que.isEmpty()){
            int size= que.size();
            for(int i=0;i<size;i++){
                TreeNode tmp = que.poll();
                dfs(tmp,targetSum,0);
                if(tmp.left!=null){
                    que.add(tmp.left);
                }
                if(tmp.right!=null){
                    que.add(tmp.right);
                }
            }
        }
        return cnt;
    }

    int cnt=0;
    public void dfs(TreeNode root,int targetSum,long sum){
        if(root==null){
            return;
        }
        sum+=root.val;
        if(sum==targetSum){
            cnt++;
        }
        dfs(root.left,targetSum,sum);
        dfs(root.right,targetSum,sum);
        sum-=root.val;
    }
}