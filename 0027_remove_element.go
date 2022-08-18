// https://leetcode.com/problems/remove-element
//
// 0027. Remove Element [Easy]
//
// Given an integer array nums and
// an integer val, remove all occurrences of val in nums in-place. The relative
// order of the elements may be changed.
//
// Since it is impossible to change the length of the array in some languages,
// you must instead have the answer be placed in the first part of the array
// nums. More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final answer. It does not
// matter what you leave beyond the first k elements.
//
// Return k after placing the final answer in the first k slots of nums.
//
// Do not allocate extra space for another array. You must do this by modifying
// the input array in-place with O(1) extra memory.
//
// Constraints:
//
//	0 <= nums.length <= 100
//	0 <= nums[i] <= 50
//	0 <= val <= 100

package leetcode

func RemoveElement(numbers []int, value int) int {
	i, j := 0, len(numbers)

	for i < j {
		if numbers[i] == value {
			numbers[i] = numbers[j-1]
			j--
		} else {
			i++
		}
	}

	return j
}
