class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String,List<String>>m = new HashMap<>();
        for(String s:strs){
            char[] sortedS = s.toCharArray();
            Arrays.sort(sortedS);
            //如果keynew String(sortedS)不存在，则创建一个新的ArrayList并将其放入map中
            //如果key存在，则将s添加到对应的ArrayList中
            m.computeIfAbsent(new String(sortedS),_->new ArrayList<>()).add(s);
        }
        //将map中的所有值转换为List<List<String>>并返回
        //使用new ArrayList<>(m.values())将map的值转换为List
        return new ArrayList<>(m.values());
    }
}