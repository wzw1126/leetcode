import java.util.HashMap;
import java.util.Map;

class Solution {
    public Node copyRandomList(Node head) {
        if (head == null) return null;

        // 映射：原节点 -> 复制节点
        Map<Node, Node> map = new HashMap<>();

        // 第一遍：创建所有新节点（只复制 val），并放入映射
        Node cur = head;
        while (cur != null) {
            map.put(cur, new Node(cur.val, null, null));
            cur = cur.next;
        }

        // 第二遍：根据映射设置 next 和 random 指针
        cur = head;
        while (cur != null) {
            Node copy = map.get(cur);
            copy.next = map.get(cur.next);       // 如果 cur.next 为 null，则 map.get 返回 null
            copy.random = map.get(cur.random);   // 同理处理 random
            cur = cur.next;
        }

        // 返回头结点的复制
        return map.get(head);
    }
}