// https://leetcode.com/problems/remove-duplicates-from-sorted-list
//
// 0083. Remove Duplicates from Sorted List [Easy]
//
// Given the head of a sorted linked list, delete all duplicates such that each
// element appears only once. Return the linked list sorted as well.
//
// Constraints:
//
//    The number of nodes in the list is in the range [0, 300].
//    -100 <= Node.val <= 100
//    The list is guaranteed to be sorted in ascending order.

package leetcode

import "leetcode/linkedlist"

func DeleteDuplicates(list *linkedlist.Node[int]) *linkedlist.Node[int] {
	if list == nil || list.Next == nil {
		return list
	}

	head := list
	for prev, curr := list, list.Next; curr != nil; curr = curr.Next {
		if curr.Value == prev.Value {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}

	return head
}
