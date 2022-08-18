// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box
//
// 1769. Minimum Number of Operations to Move All Balls to Each Box [Medium]
//
// You have n boxes. You are given a binary string boxes of length n, where
// boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.
//
// In one operation, you can move one ball from a box to an adjacent box. Box i
// is adjacent to box j if abs(i - j) == 1. Note that after doing so, there may
// be more than one ball in some boxes.
//
// Return an array answer of size n, where answer[i] is the minimum number of
// operations needed to move all the balls to the ith box.
//
// Each answer[i] is calculated considering the initial state of the boxes.
//
// Constraints:
//
//	n == boxes.length
//	1 <= n <= 2000
//	boxes[i] is either '0' or '1'.

package leetcode

func MinOperations(boxes string) []int {
	result := make([]int, len(boxes))

	// 1. from left to right, for each box count balls to the left:
	balls, moves := 0, 0
	for i := 0; i < len(boxes); i++ {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	// 2. from right to left, for each box count balls to the right:
	balls, moves = 0, 0
	for i := len(boxes) - 1; i >= 0; i-- {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	return result
}
