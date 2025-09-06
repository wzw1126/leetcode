func hammingWeight(n int) int {
    cnt:=0
    for n>0{
        cnt+=n&1
        n=n>>1
    }
    return cnt
}