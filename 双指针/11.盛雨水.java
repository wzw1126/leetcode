class Solution {
    public int maxArea(int[] height) {
        int i=0,j=height.length-1;
        int res=0;
        while(i<j){
            int h=0;
            if(height[i]<height[j]){
                h=height[i];
                i++;
            }else{
                h=height[j];
                j--;
            }
            int s=(j-i+1)*h;
            res=Math.max(res,s);
        }
        return res;
    }
}