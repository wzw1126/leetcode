class MinStack {
    Stack<Integer> st;
    Stack<Integer> min_stack;
    public MinStack() {
        st = new Stack();
        min_stack = new Stack();
    }
    
    public void push(int val) {
        st.push(val);
        if(min_stack.isEmpty() || val<=min_stack.peek()){
            min_stack.push(val);
        }
    }
    
    public void pop() {
        if(st.pop().equals(min_stack.peek())){
            min_stack.pop();
        }
    }
    
    public int top() {
        return st.peek();
    }
    
    public int getMin() {
        return min_stack.peek();    
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(val);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */