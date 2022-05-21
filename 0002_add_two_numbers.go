// https://leetcode.com/problems/add-two-numbers
package leetcode

import "github.com/junk1tm/leetcode/linkedlist"

// 0002. Add Two Numbers [Medium]
//
// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit. Add the two numbers and return the sum as a linked
// list.
//
// You may assume the two numbers do not contain any leading zero, except the
// number 0 itself.
//
// Constraints:
//
//    The number of nodes in each linked list is in the range [1, 100].
//    0 <= Node.val <= 9
//    It is guaranteed that the list represents a number that does not have leading zeros.
//
func AddTwoNumbers(list1, list2 *linkedlist.Node[int]) *linkedlist.Node[int] {
	list3 := new(linkedlist.Node[int])
	temp := list3

	overflow := false
	for list1 != nil || list2 != nil {
		list3.Next = new(linkedlist.Node[int])
		list3 = list3.Next
		if list1 != nil {
			list3.Value += list1.Value
			list1 = list1.Next
		}
		if list2 != nil {
			list3.Value += list2.Value
			list2 = list2.Next
		}
		if overflow {
			list3.Value++
			overflow = false
		}
		if list3.Value > 9 {
			list3.Value -= 10
			overflow = true
		}
	}
	if overflow {
		list3.Next = &linkedlist.Node[int]{Value: 1, Next: nil}
	}

	return temp.Next
}
