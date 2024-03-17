package valid_parenthesis_20

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		// get last element of the stack
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		last := stack[len(stack)-1]
		// check if last is the closing element of current element
		// mapping[c], c = "(" -> ")"
		if last == mapping[c] {
			// pop the last element
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
