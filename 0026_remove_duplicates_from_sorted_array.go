// https://leetcode.com/problems/remove-duplicates-from-sorted-array
//
// 0026. Remove Duplicates from Sorted Array [Easy]
//
// Given an integer array nums sorted in non-decreasing order, remove the
// duplicates in-place such that each unique element appears only once. The
// relative order of the elements should be kept the same.
//
// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array
// nums. More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result. It does not
// matter what you leave beyond the first k elements.
//
// Return k after placing the final result in the first k slots of nums.
//
// Do not allocate extra space for another array. You must do this by modifying
// the input array in-place with O(1) extra memory.
//
// Constraints:
//
//	1 <= nums.length <= 3 * 104
//	-100 <= nums[i] <= 100
//	nums is sorted in non-decreasing order.

package leetcode

import "sort"

func RemoveDuplicates(numbers []int) int {
	const max = 100

	curr, total := numbers[0], 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == curr {
			numbers[i] = max + 1
			total++
		} else {
			curr = numbers[i]
		}
	}

	sort.Ints(numbers)
	return len(numbers) - total
}
