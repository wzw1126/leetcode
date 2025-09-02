func maxAbsoluteSum(nums []int) int {
    n:=len(nums)
    maxValue:=nums[0]
    minValue:=nums[0]
    fn1:=make([]int,n)
    fn2:=make([]int,n)
    fn1[0],fn2[0]=nums[0],nums[0]
    for i:=1;i<n;i++{
        fn1[i]=max(fn1[i-1],0)+nums[i]
        fn2[i]=min(fn2[i-1],0)+nums[i]
        maxValue=max(maxValue,fn1[i])
        minValue=min(minValue,fn2[i])
    }
    if maxValue>-minValue{
        return maxValue
    }
    return -minValue

}