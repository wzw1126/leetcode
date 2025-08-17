class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        int n=nums.length;
        for(int i=0;i<nums.length-2;i++){
            if(i>0 && nums[i-1]==nums[i]){
                continue;
            }
            if(nums[i] + nums[i+1] + nums[i+2] >0){
                break;
            }
            if(nums[i]+nums[n-1]+nums[n-2] < 0){
                continue;
            }
            for(int j=i+1,k=n-1;j<k;){
                int s=nums[i]+nums[j]+nums[k];
                if(s>0){
                    k--;
                }else if(s<0){
                    j++;
                }else{
                    List<Integer> list = new ArrayList<>();
                    list.add(nums[i]);
                    list.add(nums[j]);
                    list.add(nums[k]);
                    res.add(list);
                    while(j<k&&nums[j]==nums[j+1])j++;
                    while(j<k&&nums[k]==nums[k-1])k--;
                    j++;
                    k--;
                }
            }
        }
        return res;
    }
}