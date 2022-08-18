// https://leetcode.com/problems/merge-sorted-array
//
// 0088. Merge Sorted Array [Easy]
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing
// order, and two integers m and n, representing the number of elements in nums1
// and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be
// stored inside the array nums1. To accommodate this, nums1 has a length of m +
// n, where the first m elements denote the elements that should be merged, and
// the last n elements are set to 0 and should be ignored. nums2 has a length of
// n.
//
// Constraints:
//
//	nums1.length == m + n
//	nums2.length == n
//	0 <= m, n <= 200
//	1 <= m + n <= 200
//	-109 <= nums1[i], nums2[j] <= 109

package leetcode

func MergeSortedArrays(numbers1, numbers2 []int, m, n int) {
	for m > 0 && n > 0 {
		if numbers1[m-1] > numbers2[n-1] {
			numbers1[m+n-1] = numbers1[m-1]
			m--
		} else {
			numbers1[m+n-1] = numbers2[n-1]
			n--
		}
	}
	if n > 0 {
		copy(numbers1[:n], numbers2[:n])
	}
}
