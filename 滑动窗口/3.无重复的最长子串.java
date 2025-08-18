class Solution {
  public int lengthOfLongestSubstring(String s) {
    char[] strs = s.toCharArray();
    Map<Character,Integer>map = new HashMap<>();
    int l=0,r=0,res=0;
    for(;r<strs.length;r++){
      while(map.containsKey(strs[r])){
        map.remove(strs[l]);
        l++;
      }
      map.put(strs[r],1);
      res=Math.max(res,r-l+1);
    }
    return res;

  }
}

