// https://leetcode.com/problems/longest-substring-without-repeating-characters
package leetcode

// 0003. Longest Substring Without Repeating Characters [Medium]
//
// Given a string s, find the length of the longest substring without repeating
// characters.
//
// Constraints:
//
//    0 <= s.length <= 5 * 104
//    s consists of English letters, digits, symbols and spaces.
//
func LengthOfLongestSubstring(s string) int {
	var longest int
	chars := make(map[rune]int) // char:index

	var start int // the left border of the sliding window.
	for index, char := range []rune(s) {
		if i, ok := chars[char]; ok && i >= start {
			// we've got a duplicate:
			// move the left border to the char after the duplicate.
			start = i + 1
		}
		if current := index - start + 1; current > longest {
			longest = current
		}
		chars[char] = index
	}

	return longest
}
