package brackets

func isMatch(a, b byte) bool {
	match := string([]byte{a, b})
	switch match {
	case "()", "[]", "{}":
		return true
	default:
		return false
	}
}

func parseBrackets(b byte) (open bool, ok bool) {
	switch b {
	case '(', '{', '[':
		return true, true
	case ')', '}', ']':
		return false, true
	default:
		return false, false
	}
}

func isBalanced(str string) bool {
	var stack []byte
	for _, val := range []byte(str) {
		open, ok := parseBrackets(val)
		if !ok {
			continue
		}
		if open {
			stack = append(stack, val)
			continue
		}

		if len(stack) == 0 {
			return false
		}
		if !isMatch(stack[len(stack)-1], val) {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}
