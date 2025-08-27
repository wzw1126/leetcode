func decodeString(s string) string {
	stackNum := []int{}    // 数字栈
	stackStr := []string{} // 字符串栈

	num := 0  // 当前累积数字
	cur := "" // 当前累积字符串

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch >= '0' && ch <= '9':
			// 多位数累积
			num = num*10 + int(ch-'0')

		case ch == '[':
			// 同时把当前串与数字入栈，重置
			stackNum = append(stackNum, num)
			stackStr = append(stackStr, cur)
			num = 0
			cur = ""

		case ch == ']':
			// 出栈：把当前cur重复k次，接到之前的prev后面
			k := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]

			prev := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]

			cur = prev + strings.Repeat(cur, k)

		default:
			// 这里是字母
			cur += string(ch)
		}
	}
	return cur
}