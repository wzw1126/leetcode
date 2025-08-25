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
    public TreeNode sortedArrayToBST(int[] nums) {
        return dfs(nums,0,nums.length);
    }

    TreeNode dfs(int[] nums,int left ,int right){
        if(left== right){
            return null;
        }
        int m=(left+right)>>>1;
        return new TreeNode(nums[m],dfs(nums,left,m),dfs(nums,m+1,right));
    }
}