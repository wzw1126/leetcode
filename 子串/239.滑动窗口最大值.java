class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int n = nums.length;
        int[] ans = new int[n-k+1];
        Deque<Integer> dq = new ArrayDeque<>();
        for(int i=0;i<n;i++){
            // 1. 右边入
            while(!dq.isEmpty() && nums[dq.getLast()]<=nums[i]){
                dq.removeLast(); // 维护 dq 的单调性
            }
            dq.addLast(i);
            // 2. 左边出
            int left = i-k+1;
            if(dq.getFirst()<left){
                dq.removeFirst();
            }
            // 3. 在窗口左端点处记录答案
            if(left>=0){
                // 由于队首到队尾单调递减，所以窗口最大值就在队首
                ans[left]=nums[dq.getFirst()];
            }
        }
        return ans;
    }
}