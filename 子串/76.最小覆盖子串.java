class Solution {
    public String minWindow(String s, String t) {
        if(t.isEmpty() || s.length()<t.length()) return "";
        Map<Character,Integer> need = new HashMap<>();
        for(char ch : t.toCharArray()){
            need.put(ch,need.getOrDefault(ch,0)+1);
        }
        Map<Character,Integer> win = new HashMap<>();
        int valid=0,left=0,start=0,minLen= Integer.MAX_VALUE;

        for(int right=0;right<s.length();right++){
            char c = s.charAt(right);
            if(need.containsKey(c)){
                win.put(c,win.getOrDefault(c,0)+1);
                if(win.get(c).intValue() == need.get(c).intValue()) valid++;
            }
            while(valid == need.size()){
                if(right-left+1 < minLen){
                    minLen=right-left+1;
                    start=left;
                }
                char d = s.charAt(left++);
                if(need.containsKey(d)){
                    if(win.get(d).intValue() == need.get(d).intValue()) valid--;
                    win.put(d,win.get(d)-1);
                }
            }
        }
        return minLen == Integer.MAX_VALUE? "":s.substring(start,start+minLen);
    }
}