func nextPermutation(nums []int)  {

    /*
      1.从右往左找第一个“下降位”索引 i：nums[i] < nums[i+1]（注意返回的是 i 而不是 i+1）。
      2.在右侧尾部找到比 nums[i] 大的最小元素（从右往左第一个 > nums[i] 的位置 j），交换 nums[i] 和 nums[j].
	  3.再把 i+1..end 整段反转（使之最小化）。
    */

    n:=len(nums)
    i:=n-2
    for i>=0{
        if nums[i+1]>nums[i]{
            break
        }
        i--
    }
    if i>=0{
        j:=n-1
        for j>=0 && nums[j]<=nums[i]{
            j--
        }
        nums[i],nums[j] = nums[j],nums[i]
    }
    for l,r:=i+1,n-1;l<r;{
        nums[l],nums[r]=nums[r],nums[l]
        l++
        r--
    }
}