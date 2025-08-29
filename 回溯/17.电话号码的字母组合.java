class Solution {
    List<String> res;
    public List<String> letterCombinations(String digits) {
        res = new ArrayList<>();
        if(digits == null || digits.length()==0){
            return res;
        }
        String[] numString ={"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"};
        backing(digits,numString,0);
        return res;
    }
    StringBuilder sb = new StringBuilder();
    void backing(String digits,String[] numString,int start){
        if(start == digits.length()){
            res.add(sb.toString());
            return;
        }
        String str = numString[digits.charAt(start)-'0'];
        for(int i=0;i<str.length();i++){
            sb.append(str.charAt(i));
            backing(digits,numString,start+1);
            sb.deleteCharAt(sb.length()-1);
        }
    }
}