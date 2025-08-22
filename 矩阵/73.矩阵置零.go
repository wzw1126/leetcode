package main
func setZeroes(matrix [][]int)  {
    m,n:=len(matrix),len(matrix[0])
    arri:=make([]bool,m)
    arrj:=make([]bool,n)
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if matrix[i][j]==0{
                arri[i]=true
                arrj[j]=true
            }
        }
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if arri[i] || arrj[j]{
                matrix[i][j]=0
            }
            
        }
    }
}