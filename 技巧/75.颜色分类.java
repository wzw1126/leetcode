class Solution {
    public void sortColors(int[] nums) {
        sort(nums,0,nums.length-1);
    }

    void  sort(int[] nums,int left,int right){
        if(left<right){
            int pivotIndex = solu(nums,left,right);
            sort(nums,left,pivotIndex-1);
            sort(nums,pivotIndex+1,right);
        }
    }
    int solu(int[] nums,int left,int right){
        int pivot=nums[left];
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