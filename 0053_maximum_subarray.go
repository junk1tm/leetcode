// https://leetcode.com/problems/maximum-subarray
package leetcode

// 0053. Maximum Subarray [Easy]
//
// Given an integer array nums, find the contiguous subarray (containing at
// least one number) which has the largest sum and return its sum.
//
// A subarray is a contiguous part of an array.
//
// Constraints:
//
//    1 <= nums.length <= 105
//    -104 <= nums[i] <= 104
//
func MaxSubArray(numbers []int) int {
	sum, max := numbers[0], numbers[0]
	for _, n := range numbers[1:] {
		if sum < 0 {
			sum = n
		} else {
			sum += n
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
