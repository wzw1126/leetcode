//go当中就不能使用x+""了，只能使用strconv.Itoa(x)
class Solution {
    public boolean isPalindrome(int x) {
        if(x<0){
            return false;
        }
        
        String s=x+"";
        int i=0,j=s.length()-1;
        while(i<j){
            if(s.charAt(i++)!=s.charAt(j--)){
                return false;
            }
        }
        return true;
    }
}