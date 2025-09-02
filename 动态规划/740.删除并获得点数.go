//值域的打家劫舍
//首先先求出对应的值的总和，和最大的key，然后再进行打家劫舍

func deleteAndEarn(nums []int) int {
    m:=make(map[int]int,0)
    maxValue:=0
    for i:=0;i<len(nums);i++{
        m[nums[i]]=m[nums[i]]+nums[i]
        maxValue=max(maxValue,nums[i])
    }
    arr:=make([]int,maxValue+1)
    arr[1]=m[1]
    for i:=2;i<=maxValue;i++{
        arr[i]=max(arr[i-2]+m[i],arr[i-1])
    }
    return arr[maxValue]
    
}