func plusOne(digits []int) []int {
    cnt:=1
    n:=len(digits)
    for i:=n-1;i>=0 && cnt>0;i--{
        cnt+=digits[i]
        digits[i]=cnt%10
        cnt/=10
    }
    if cnt>0{
        t:=make([]int,n+1)
        t[0]=cnt
        copy(t[1:],digits)
        return t
    }else{
        return digits
    }
    
}