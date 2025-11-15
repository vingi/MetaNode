package task1

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func ValidParentheses(s string) bool {
	// 使用切片作为栈
	var stack []rune

	// 遍历字符串中的每个字符
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 左括号：压入栈中
			stack = append(stack, char)
		case ')', '}', ']':
			// 右括号：检查栈顶是否是对应的左括号
			if len(stack) == 0 {
				// 栈为空，没有匹配的左括号
				return false
			}

			// 检查栈顶元素是否与当前右括号匹配
			top := stack[len(stack)-1]
			if !isMatchingPair(top, char) {
				return false
			}

			// 移出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	// 遍历结束后，栈应该为空才表示所有括号都正确匹配
	return len(stack) == 0
}

// isMatchingPair 检查两个括号是否匹配
func isMatchingPair(left, right rune) bool {
	switch left {
	case '(':
		return right == ')'
	case '{':
		return right == '}'
	case '[':
		return right == ']'
	default:
		return false
	}
}
