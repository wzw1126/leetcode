/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //层次遍历，每次只选择最右边的节点
func rightSideView(root *TreeNode) []int {
    res:=[]int{}
    if root==nil{
        return res
    }
    que:=[]*TreeNode{root}
    for len(que)>0{
        n:=len(que)
        for i:=0;i<n;i++{
            tmp:=que[i]
            if tmp.Left!=nil{
                que=append(que,tmp.Left)
            }
            if tmp.Right!=nil{
                que=append(que,tmp.Right)
            }
            if i==n-1{
                res=append(res,tmp.Val)
            }
        }
        que=que[n:]
    }
    return res
}

//先右节点，每次达到层数就第一个加进去
func rightSideView(root *TreeNode) []int {
    
    var dfs func(*TreeNode,int)
    res:=[]int{}
    dfs = func(root *TreeNode,depth int){
        if root==nil{
            return
        }
        if depth == len(res){
            res=append(res,root.Val)
        }
        dfs(root.Right,depth+1)
        dfs(root.Left,depth+1)
    }
    dfs(root,0)
    return res
}