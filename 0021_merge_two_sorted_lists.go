// https://leetcode.com/problems/merge-two-sorted-lists
package leetcode

import "github.com/junk1tm/leetcode/linkedlist"

// 0021. Merge Two Sorted Lists [Easy]
//
// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists in a one sorted list. The list should be made by splicing
// together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Constraints:
//
//    The number of nodes in both lists is in the range [0, 50].
//    -100 <= Node.val <= 100
//    Both list1 and list2 are sorted in non-decreasing order.
//
func MergeTwoLists(list1, list2 *linkedlist.Node[int]) *linkedlist.Node[int] {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	list3 := new(linkedlist.Node[int])
	temp := list3

	for list1 != nil && list2 != nil {
		list3.Next = new(linkedlist.Node[int])
		list3 = list3.Next
		if list1.Value < list2.Value {
			list3.Value = list1.Value
			list1 = list1.Next
		} else {
			list3.Value = list2.Value
			list2 = list2.Next
		}
	}

	if list1 != nil {
		list3.Next = list1
	} else {
		list3.Next = list2
	}

	return temp.Next
}
