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
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        int n = preorder.length;
        if(n==0){
            return null;
        }
        int leftSize = findIndex(inorder,preorder[0]);
        int[] pre1 = Arrays.copyOfRange(preorder,1,1+leftSize);
        int[] pre2 = Arrays.copyOfRange(preorder,1+leftSize,n);
        int[] ino1 = Arrays.copyOfRange(inorder,0,leftSize);
        int[] ino2 = Arrays.copyOfRange(inorder,leftSize+1,n);
        TreeNode left = buildTree(pre1,ino1);
        TreeNode right = buildTree(pre2,ino2);
        return new TreeNode(preorder[0],left,right);

    }

    int findIndex(int[] inorder,int x){
        for(int i=0;;i++){
            if(inorder[i]==x){
                return i;
            }
        }
    }
}