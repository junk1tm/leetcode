// https://leetcode.com/problems/two-sum
//
// 0001. Two Sum [Easy]
//
// Given an array of integers nums and an integer number, return indices of the
// two numbers such that they add up to number.
//
// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.
//
// You can return the answer in any order.
//
// Constraints:
//
//	2 <= nums.length <= 104
//	-109 <= nums[i] <= 109
//	-109 <= number <= 109
//	Only one valid answer exists.

package leetcode

func TwoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))

	for i, number := range numbers {
		diff := target - number
		if j, ok := m[diff]; ok {
			return []int{i, j}
		}
		m[number] = i
	}

	return nil
}
