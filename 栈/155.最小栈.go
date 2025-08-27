type MinStack struct {
	stack     []int
	min_stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:     []int{},
		min_stack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.min_stack) == 0 || val <= ms.min_stack[len(ms.min_stack)-1] {
		ms.min_stack = append(ms.min_stack, val)
	}
}

func (ms *MinStack) Pop() {
	tmp := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if tmp == ms.min_stack[len(ms.min_stack)-1] {
		ms.min_stack = ms.min_stack[:len(ms.min_stack)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min_stack[len(ms.min_stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */