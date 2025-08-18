class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        int n = p.length();
        if(s.length()<n) return new ArrayList<Integer>();
        Map<Character,Integer>mapp = new HashMap<>();
        for(Character ch:p.toCharArray()){
            mapp.put(ch,mapp.getOrDefault(ch,0)+1);
        }
        Map<Character,Integer>maps = new HashMap<>();
        for(int i=0;i<n;i++){
            maps.put(s.charAt(i),maps.getOrDefault(s.charAt(i),0)+1);
        }
        List<Integer> list = new ArrayList<>();
        if(isTrue(mapp,maps)){
            list.add(0);
        }
        for(int i=n;i<s.length();i++){
            char ch1 = s.charAt(i);
            char ch2 = s.charAt(i-n);
            maps.put(ch2,maps.get(ch2)-1);
            if(maps.get(ch2)==0){
                maps.remove(ch2);
            }
            maps.put(ch1,maps.getOrDefault(ch1,0)+1);
            if(isTrue(mapp,maps)){
                list.add(i-n+1);
            }
        }
        return list;
    }
    //注意当Integer的值超过127的时候就是对象了，就不能使用!=
    boolean isTrue(Map<Character,Integer>mapp,Map<Character,Integer>maps){
        if(mapp.size()!=maps.size()) return false;
        for(Character ch:mapp.keySet()){
            if(!Objects.equals(mapp.get(ch),maps.get(ch))){
                return false;
            }
        }
        return true;
    }
}