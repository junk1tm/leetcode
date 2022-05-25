// https://leetcode.com/problems/search-insert-position
package leetcode

// 0035. Search Insert Position [Easy]
//
// Given a sorted array of distinct integers and a target value, return the
// index if the target is found. If not, return the index where it would be if
// it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Constraints:
//
//    1 <= nums.length <= 104
//    -104 <= nums[i] <= 104
//    nums contains distinct values sorted in ascending order.
//    -104 <= target <= 104
//
func SearchInsert(numbers []int, target int) int {
	var k, middle int
	for i, j := 0, len(numbers)-1; i <= j; {
		k = (i + j) / 2
		middle = numbers[k]
		switch {
		case target == middle:
			return k
		case target > middle:
			i = k + 1
		case target < middle:
			j = k - 1
		}
	}
	if target > middle {
		return k + 1
	}
	return k
}
