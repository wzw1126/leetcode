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
    int ans=0;
    public int diameterOfBinaryTree(TreeNode root) {
        dfs(root);
        return ans;
    }
    int dfs(TreeNode node){
        if(node==null){
            return -1;
        }
        int lLen = dfs(node.left)+1;
        int rLen = dfs(node.right)+1;
        ans = Math.max(ans,lLen+rLen);
        return Math.max(lLen,rLen);
    }
}