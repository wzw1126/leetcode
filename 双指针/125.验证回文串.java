class Solution {
    public boolean isPalindrome(String s) {
        StringBuilder sb = new StringBuilder();
        for(int i=0;i<s.length();i++){
            char c = Character.toLowerCase(s.charAt(i));
            if((c>='a' && c<='z') || (c>='0' && c<='9')){
                sb.append(c);
            }
        }
        String str = sb.toString();
        return str.equals(sb.reverse().toString());
    }
}