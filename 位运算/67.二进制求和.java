class Solution {
    public String addBinary(String a, String b) {
        StringBuilder sb = new StringBuilder();
        int lena = a.length();
        int lenb = b.length();
        int i=lena-1,j=lenb-1;
        int cnt=0;
        while(i>=0 || j>=0 || cnt>0){
            if(i>=0){
                cnt+=a.charAt(i)-'0';
                i--;
            }
            if(j>=0){
                cnt+=b.charAt(j)-'0';
                j--;
            }
            sb.append(cnt%2);
            cnt/=2;
        }
        return sb.reverse().toString();
    }
}