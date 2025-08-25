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
    public boolean isSymmetric(TreeNode root) {
        return isFlag(root.left,root.right);
    }
    boolean isFlag(TreeNode left,TreeNode right){
        if(left==null || right==null){
            return left==right;
        }
        return left.val==right.val && isFlag(left.left,right.right) && isFlag(left.right,right.left);
    }
}