// https://leetcode.com/problems/sqrtx
package leetcode

// 0069. Sqrt(x) [Easy]
//
// Given a non-negative integer x, compute and return the square root of x.
//
// Since the return type is an integer, the decimal digits are truncated, and
// only the integer part of the result is returned.
//
// Note: You are not allowed to use any built-in exponent function or operator,
// such as pow(x, 0.5) or x ** 0.5.
//
// Constraints:
//
//    0 <= x <= 2^31 - 1
//
func Sqrt(square int) int {
	pow := func(n int) int { return n * n }

	low, high := 0, square
	for low <= high {
		middle := (high + low) / 2
		switch {
		case pow(middle) <= square && square < pow(middle+1):
			return middle
		case square > pow(middle):
			low = middle + 1
		case square < pow(middle):
			high = middle - 1
		}
	}

	return low
}
