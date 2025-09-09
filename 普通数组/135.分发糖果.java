//思想：贪心+数学
//从左到右扫描ratings数组，找到每一个山峰，然后计算出这个山峰对应的糖果数
class Solution {
    public int candy(int[] ratings) {
        int n = ratings.length;
        int res=n;
        for(int i=0;i<n;i++){
            int start = i;
            if(i>0 && ratings[i-1]<ratings[i]){
                start--;
            }
            for(;i+1<n && ratings[i]<ratings[i+1];){
                i++;
            }
            int top=i;
            for(;i+1<n && ratings[i]>ratings[i+1];){
                i++;
            }
            int inc = top-start;
            int dec = i-top;
            res+=(inc*(inc-1)+dec*(dec-1))/2+Math.max(inc,dec);
        }
        return res;
    }
}