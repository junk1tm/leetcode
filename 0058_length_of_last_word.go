// https://leetcode.com/problems/length-of-last-word
//
// 0058. Length of Last Word [Easy]
//
// Given a string s consisting of words and spaces, return the length of the
// last word in the string.
//
// A word is a maximal substring consisting of non-space characters only.
//
// Constraints:
//
//	1 <= s.length <= 104
//	s consists of only English letters and spaces ' '.
//	There will be at least one word in s.

package leetcode

func LengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	var length int
	for i >= 0 && s[i] != ' ' {
		i--
		length++
	}

	return length
}
