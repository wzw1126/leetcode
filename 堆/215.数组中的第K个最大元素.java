//使用优先级队列
class Solution {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> minHeap = new PriorityQueue<>(k,Comparator.comparingInt(a->a));
        for(int i=0;i<k;i++){
            minHeap.offer(nums[i]);
        }
        int n = nums.length;
        for(int i=k;i<n;i++){
            if(nums[i] > minHeap.peek()){
                minHeap.poll();
                minHeap.offer(nums[i]);
            }
        }
        return minHeap.peek();
    }
}

//使用快速选择
class Solution {
    public int findKthLargest(int[] nums, int k) {
        return sort(nums,0,nums.length-1,nums.length-k);
    }
    int sort (int[]nums,int left,int right,int x){
        int index=partition(nums,left,right);
        if(index==x){
            return nums[index];
        }else if(index>x){
            return sort(nums,left,index-1,x);
        }else{
            return sort(nums,index+1,right,x);
        }

    }
    int partition(int[]nums,int left,int right){
        int pivot = nums[left];
        int i=left,j=right;
        while(i<j){
            while(i<j && nums[j]>pivot){
                j--;
            }
            while(i<j && nums[i]<=pivot){
                i++;
            }
            swap(nums,i,j);
        }
        swap(nums,i,left);
        return i;
    }

    void swap(int[] nums,int i,int j){
        int temp = nums[i];
        nums[i]=nums[j];
        nums[j]=temp;
    }
}