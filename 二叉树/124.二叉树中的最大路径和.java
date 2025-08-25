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
    Integer res = Integer.MIN_VALUE;
    public int maxPathSum(TreeNode root) {
        maxPath(root);
        return res;
    }
    int maxPath(TreeNode root){
        if(root==null){
            return 0;
        }
        int left = maxPath(root.left);
        int right = maxPath(root.right);
        res = Math.max(res,left+right+root.val);
        return Math.max(Math.max(left,right)+root.val,0);
    }
}