func canJump(nums []int) bool {
    n:=len(nums)-1
    res:=0
    for i:=0;i<=res;i++{
        res=max(res,i+nums[i])
        if res>=n{
            return true
        }
    }
    return false
}