// https://leetcode.com/problems/implement-strstr/
//
// 0028. Implement strStr() [Easy]
//
// Given two strings needle and haystack, return the index of the first
// occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
// Clarification:
//
// What should we return when needle is an empty string? This is a great
// question to ask during an interview.
//
// For the purpose of this problem, we will return 0 when needle is an empty
// string. This is consistent to C's strstr() and Java's indexOf().
//
// Constraints:
//
//	1 <= haystack.length, needle.length <= 104
//	haystack and needle consist of only lowercase English characters.

package leetcode

func IndexOf(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := range haystack[:len(haystack)-len(needle)+1] {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
