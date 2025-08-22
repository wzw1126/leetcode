// merge 合并重叠区间，返回一个不重叠且覆盖所有输入区间的区间数组。
// 思路：先按区间左端点排序，然后线性扫描，维护当前待合并区间 cur。
// 遇到下一个区间若与 cur 重叠（next.start <= cur.end），则更新 cur.end = max(cur.end, next.end)；
// 否则将 cur 入结果并把 cur 更新为下一个区间。
// 时间复杂度：O(n log n)（排序），空间复杂度：O(n)（用于结果）。
class Solution {
    public int[][] merge(int[][] intervals) {
        if(intervals==null || intervals.length == 0){
            return new int[0][];
        }
        Arrays.sort(intervals,(a,b)->{
            if(a[0] == b[0]){
                return Integer.compare(a[1],b[1]);
            }
            return Integer.compare(a[0],b[0]);
        });
        List<int[]> list = new ArrayList<>();
        int[] res = new int[2];
        res[0]=intervals[0][0];
        res[1]=intervals[0][1];
        for(int i=1;i<intervals.length;i++){
            if(intervals[i][0]<=res[1]){
                res[1]=Math.max(intervals[i][1],res[1]);
            }else{
                list.add(res);
                res=new int[2];
                res[0]=intervals[i][0];
                res[1]=intervals[i][1];
            }
        }
        list.add(res);
        return list.toArray(new int[list.size()][]);
    }
}