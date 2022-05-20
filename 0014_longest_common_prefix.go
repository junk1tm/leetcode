// https://leetcode.com/problems/longest-common-prefix
package leetcode

// 0014. Longest Common Prefix [Easy]
//
// Write a function to find the longest common prefix string amongst an array of
// strings.
//
// If there is no common prefix, return an empty string "".
//
// Constraints:
//
//    1 <= strs.length <= 200
//    0 <= strs[i].length <= 200
//    strs[i] consists of only lower-case English letters.
//
func LongestCommonPrefix(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	var i int

loop:
	for {
		set := make(map[byte]struct{})
		for _, s := range strings {
			if i >= len(s) {
				break loop
			}
			set[s[i]] = struct{}{}
		}
		if len(set) > 1 {
			break loop
		}
		i++
	}

	return strings[0][:i]
}
