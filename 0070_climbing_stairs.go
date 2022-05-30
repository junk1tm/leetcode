// https://leetcode.com/problems/climbing-stairs
package leetcode

// 0070. Climbing Stairs [Easy]
//
// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can
// you climb to the top?
//
// Constraints:
//
//    1 <= n <= 45
//
func ClimbingStairs(steps int) int {
	switch steps {
	case 1:
		return 1
	case 2:
		return 2
	}

	paths := make([]int, steps)
	paths[0] = 1
	paths[1] = 2
	for i := 2; i < steps; i++ {
		paths[i] = paths[i-1] + paths[i-2]
	}

	return paths[steps-1]
}
