// https://leetcode.com/problems/palindrome-number
//
// 0009. Palindrome Number [Easy]
//
// Given an integer x, return true if x is palindrome integer.
//
// An integer is a palindrome when it reads the same backward as forward.
//
// Constraints:
//
//	-231 <= x <= 231 - 1

package leetcode

func IsPalindrome(number int) bool {
	if number < 0 {
		return false
	}

	original, x := number, 0
	for number > 0 {
		last := number % 10
		x = x*10 + last
		number /= 10
	}

	return original == x
}
