
class Solution {    
    public int trap(int[] h) {
        int n = h.length, ans = 0;
        int left = 0, right = n - 1;
        int leftMax = 0, rightMax = 0;
        while (left < right) {
            leftMax = Math.max(leftMax, h[left]);
            rightMax = Math.max(rightMax, h[right]);
            if (leftMax < rightMax) {
                ans += leftMax - h[left];
                left++;
            } else {
                ans += rightMax - h[right];
                right--;
            }
        }
        return ans;
    }
}

//前缀和方法
class Solution {
    public int trap(int[] height) {
        int n = height.length;
        //构建前缀最大
        int[] preMax = new int[n];
        preMax[0]=height[0];
        for(int i = 1;i<n;i++){
            preMax[i] = Math.max(preMax[i-1],height[i]);
        }
        //构建后缀最大
        int[] sufMax = new int[n];
        sufMax[n-1]=height[n-1];
        for(int i = n-2;i>=0;i--){
            sufMax[i]=Math.max(sufMax[i+1],height[i]);
        }
        //逐格累加雨水
        int ans=0;
        for(int i =0;i<n;i++){
            ans+=Math.min(preMax[i],sufMax[i])-height[i];
        }
        return ans;
    }
}