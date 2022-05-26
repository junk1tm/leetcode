// https://leetcode.com/problems/plus-one
package leetcode

// 0066. Plus One [Easy]
//
// You are given a large integer represented as an integer array digits, where
// each digits[i] is the ith digit of the integer. The digits are ordered from
// most significant to least significant in left-to-right order. The large
// integer does not contain any leading 0's.
//
// Increment the large integer by one and return the resulting array of digits.
//
// Constraints:
//
//    1 <= digits.length <= 100
//    0 <= digits[i] <= 9
//    digits does not contain any leading 0's.
//
func PlusOne(digits []int) []int {
	overflow := 1 // start with +1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += overflow
		overflow, digits[i] = digits[i]/10, digits[i]%10
	}
	if overflow > 0 {
		return append([]int{overflow}, digits...)
	}
	return digits
}
