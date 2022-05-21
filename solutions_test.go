package leetcode_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/junk1tm/leetcode"
)

// 0001. Two Sum [Easy]
func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		indices []int
	}{
		{numbers: []int{2, 7, 11, 15}, target: 9, indices: []int{1, 0}},
		{numbers: []int{3, 2, 4}, target: 6, indices: []int{2, 1}},
		{numbers: []int{3, 3}, target: 6, indices: []int{1, 0}},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			indices := leetcode.TwoSum(tt.numbers, tt.target)
			equal[E](t, indices, tt.indices)
		})
	}
}

// 0009. Palindrome Number [Easy]
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		number int
		answer bool
	}{
		{number: 121, answer: true},
		{number: -121, answer: false},
		{number: 10, answer: false},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			result := leetcode.IsPalindrome(tt.number)
			equal[E](t, result, tt.answer)
		})
	}
}

// 0013. Roman to Integer [Easy]
func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman   string
		integer int
	}{
		{roman: "III", integer: 3},
		{roman: "LVIII", integer: 58},
		{roman: "MCMXCIV", integer: 1994},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			integer := leetcode.RomanToInt(tt.roman)
			equal[E](t, integer, tt.integer)
		})
	}
}

// 0014. Longest Common Prefix [Easy]
func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strings []string
		prefix  string
	}{
		{strings: []string{"flower", "flow", "flight"}, prefix: "fl"},
		{strings: []string{"dog", "racecar", "car"}, prefix: ""},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			prefix := leetcode.LongestCommonPrefix(tt.strings)
			equal[E](t, prefix, tt.prefix)
		})
	}
}

// 0020. Valid Parentheses [Easy]
func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{input: "()", valid: true},
		{input: "()[]{}", valid: true},
		{input: "(]", valid: false},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			valid := leetcode.IsValidParentheses(tt.input)
			equal[E](t, valid, tt.valid)
		})
	}
}

// 0026. Remove Duplicates from Sorted Array [Easy]
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		numbers []int
		length  int
		result  []int
	}{
		{numbers: []int{1, 1, 2}, length: 2, result: []int{1, 2}},
		{numbers: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, length: 5, result: []int{0, 1, 2, 3, 4}},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			length := leetcode.RemoveDuplicates(tt.numbers)
			equal[E](t, length, tt.length)
			equal[E](t, tt.numbers[:length], tt.result)
		})
	}
}

// 0027. Remove Element [Easy]
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		numbers []int
		value   int
		length  int
		result  []int
	}{
		{numbers: []int{3, 2, 2, 3}, value: 3, length: 2, result: []int{2, 2}},
		{numbers: []int{0, 1, 2, 2, 3, 0, 4, 2}, value: 2, length: 5, result: []int{0, 1, 4, 0, 3}},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			length := leetcode.RemoveElement(tt.numbers, tt.value)
			equal[E](t, length, tt.length)
			equal[E](t, tt.numbers[:length], tt.result)
		})
	}
}

// 1769. Minimum Number of Operations to Move All Balls to Each Box [Medium]
func TestMinOperations(t *testing.T) {
	tests := []struct {
		boxes      string
		operations []int
	}{
		{boxes: "110", operations: []int{1, 1, 3}},
		{boxes: "001011", operations: []int{11, 8, 5, 4, 3, 4}},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			operations := leetcode.MinOperations(tt.boxes)
			equal[E](t, operations, tt.operations)
		})
	}
}

func name(i int) string {
	return fmt.Sprintf("example %d", i+1)
}

type (
	E *testing.T
	F *testing.T
)

func equal[T E | F](t T, got, want any) {
	(*testing.T)(t).Helper()
	if !reflect.DeepEqual(got, want) {
		if got == "" {
			got = "<empty>"
		}
		if want == "" {
			want = "<empty>"
		}
		f(t)("got %v; want %v", got, want)
	}
}

func f[T E | F](t T) func(format string, args ...any) {
	switch any(t).(type) {
	case E:
		return (*testing.T)(t).Errorf
	case F:
		return (*testing.T)(t).Fatalf
	default:
		panic("unreachable")
	}
}
