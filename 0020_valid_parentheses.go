// https://leetcode.com/problems/valid-parentheses
package leetcode

// 0020. Valid Parentheses [Easy]
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and
// ']', determine if the input string is valid.
//
// An input string is valid if:
//
//    Open brackets must be closed by the same type of brackets.
//    Open brackets must be closed in the correct order.
//
// Constraints:
//
//    1 <= s.length <= 104
//    s consists of parentheses only '()[]{}'.
//
func IsValidParentheses(input string) bool {
	if len(input)%2 != 0 {
		return false
	}

	isPair := func(b1, b2 rune) bool {
		switch {
		case b1 == '(' && b2 == ')':
			return true
		case b1 == '[' && b2 == ']':
			return true
		case b1 == '{' && b2 == '}':
			return true
		default:
			return false
		}
	}

	var stack []rune
	for _, curr := range []rune(input) {
		if curr == '(' || curr == '[' || curr == '{' {
			stack = append(stack, curr)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if prev := stack[len(stack)-1]; !isPair(prev, curr) {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
