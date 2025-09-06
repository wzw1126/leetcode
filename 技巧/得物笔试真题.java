//给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，再给你一个k，求直线上大于k个点的直线有多少条

//用一对点唯一确定一条直线，把每条直线规范化成整数三元组 (A,B,C)（表示方程 Ax + By = C），
// 并把所有在这条线上出现过的点的索引放到集合里。最后统计“点数 > k”的直线条数
import java.util.*;

class Solution {
    // 统计直线上点数 > k 的直线条数
    public int countLinesWithMoreThanK(int[][] points, int k) {
        int n = points.length;
        if (n == 0 || k >= n) return 0;

        Map<LineKey, HashSet<Integer>> line2pts = new HashMap<>();

        for (int i = 0; i < n; i++) {
            long x1 = points[i][0], y1 = points[i][1];
            for (int j = i + 1; j < n; j++) {
                long x2 = points[j][0], y2 = points[j][1];

                // 构造规范化的直线 Ax + By = C
                long A = y2 - y1;
                long B = x1 - x2;       // 注意 B = -(x2 - x1)
                long C = A * x1 + B * y1;

                LineKey key = normalize(A, B, C);

                // 把该直线上的两个点索引加入集合
                line2pts.computeIfAbsent(key, __ -> new HashSet<>()).add(i);
                line2pts.get(key).add(j);
            }
        }

        int ans = 0;
        for (HashSet<Integer> set : line2pts.values()) {
            if (set.size() > k) ans++;
        }
        return ans;
    }

    // 规范化 (A,B,C):
    // 1) 约去 g = gcd(|A|,|B|,|C|)
    // 2) 固定符号：若 A<0 或 (A==0 且 B<0) 则整体取反，保证唯一性
    private LineKey normalize(long A, long B, long C) {
        long g = gcd(abs(A), gcd(abs(B), abs(C)));
        if (g != 0) { A /= g; B /= g; C /= g; }

        // 固定符号
        if (A < 0 || (A == 0 && B < 0)) {
            A = -A; B = -B; C = -C;
        }
        return new LineKey(A, B, C);
    }

    private long abs(long x) { return x >= 0 ? x : -x; }

    private long gcd(long a, long b) {
        while (b != 0) {
            long t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    // 直线键
    static class LineKey {
        final long A, B, C;
        LineKey(long A, long B, long C) { this.A = A; this.B = B; this.C = C; }
        @Override public boolean equals(Object o) {
            if (this == o) return true;
            if (!(o instanceof LineKey)) return false;
            LineKey that = (LineKey) o;
            return A == that.A && B == that.B && C == that.C;
        }
        @Override public int hashCode() {
            int h = Long.hashCode(A);
            h = h * 31 + Long.hashCode(B);
            h = h * 31 + Long.hashCode(C);
            return h;
        }
    }
}
