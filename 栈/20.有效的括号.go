//我自己写的思路
func isValid1(s string) bool {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, ')')
			continue
		} else if s[i] == '{' {
			st = append(st, '}')
			continue
		} else if s[i] == '[' {
			st = append(st, ']')
			continue
		}
		if len(st) == 0 || st[len(st)-1] != s[i] {
			return false
		}
		st = st[:len(st)-1]
	}
	if len(st) != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	if len(s)%2 != 0 { // s 长度必须是偶数
		return false
	}
	st := []rune{}
	for _, c := range s {
		switch c {
		case '(':
			st = append(st, ')') // 入栈对应的右括号
		case '[':
			st = append(st, ']')
		case '{':
			st = append(st, '}')
		default: // c 是右括号
			if len(st) == 0 || st[len(st)-1] != c {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0 // 所有左括号必须匹配完毕
}
