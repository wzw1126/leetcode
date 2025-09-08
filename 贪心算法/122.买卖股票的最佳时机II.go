//贪心做法，只要后一天的价格比前一天高，就在前一天买入，后一天卖出
func maxProfit(prices []int) int {
    sum:=0
    tmp:=prices[0]
    for i:=1;i<len(prices);i++{
        if prices[i]>tmp{
            sum+=prices[i]-tmp
        }
        tmp=prices[i]
    }
    return sum
}