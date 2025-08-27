class Solution {
    public boolean isValid(String s) {
        if(s.length()%2 !=0){
            return false;
        }
        Deque<Character> st = new ArrayDeque<>();
        for(char ch : s.toCharArray()){
            if(ch == '('){
                st.push(')');
            }else if(ch == '['){
                st.push(']');
            }else if(ch == '{'){
                st.push('}');
            }else if(st.isEmpty() || st.pop()!=ch){
                return false;
            }
        }
        return st.isEmpty();
    }
}