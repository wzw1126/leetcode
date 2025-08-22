package 普通数组
func merge(intervals [][]int) [][]int {
    n:=len(intervals)
    if n==0 {
        return nil
    }
    sort.Slice(intervals,func(i,j int)bool{
        if intervals[i][0] == intervals[j][0]{
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0]<intervals[j][0]
    })

    res:=make([][]int,0,n)
    cur := make([]int,2)
    copy(cur,intervals[0])
    for i:=1;i<n;i++{
        next:=intervals[i]
        if next[0]<=cur[1]{
            if next[1]>cur[1]{
                cur[1]=next[1]
            }
        }else{
            res = append(res,cur)
            cur = make([]int,2)
            copy(cur,next)
        }
    }
    res = append(res,cur)
    return res
}