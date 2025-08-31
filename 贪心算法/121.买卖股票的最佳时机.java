class Solution {
    public int maxProfit(int[] prices) {
        int premin=prices[0];
        int res=0;
        for(int x:prices){
            res=Math.max(res,x-premin);
            if(x<premin){
                premin=x;
            }

        }
        return res;
    }
}