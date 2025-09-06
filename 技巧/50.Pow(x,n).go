
//https://leetcode.cn/problems/powx-n/solutions/2858114/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/?envType=study-plan-v2&envId=top-interview-150
//主要思路：
//快速幂算法的核心思想是将指数 n 分解为二进制形式，然后通过平方和乘法来高效地计算 x 的 n 次方。
//1. 处理负指数：如果 n 是负数，我们将 x 替换为 1/x，并将 n 取绝对值。这是因为 x^(-n) = 1/(x^n)。
//2. 初始化结果：我们初始化一个变量 ans 为 1，用于存储最终的结果。
//3. 循环计算：我们使用一个 while 循环，直到 n 变为 0。在每次循环中，我们检查 n 的最低位（通过 n & 1 判断）是否为 1。如果是，我们将当前的 x 乘到 ans 上，因为这表示我们需要包含这个 x 的贡献。
//4. 平方 x：无论最低位是否为 1，我们都将 x 平方（x *= x），这相
func myPow(x float64, n int) float64 {
    ans:=float64(1)
    if n<0{
        n=-n
        x=1/x
    }  
    for n>0{
        if n&1>0{
            ans*=x
        }
        x*=x
        n>>=1
    }
    return ans
}