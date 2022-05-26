package leetcode_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/junk1tm/leetcode"
	"github.com/junk1tm/leetcode/linkedlist"
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

// 0002. Add Two Numbers [Medium]
func TestAddTwoNumbers(t *testing.T) {
	listOf := linkedlist.New[int]

	tests := []struct {
		list1  *linkedlist.Node[int]
		list2  *linkedlist.Node[int]
		result *linkedlist.Node[int]
	}{
		{list1: listOf(2, 4, 3), list2: listOf(5, 6, 4), result: listOf(7, 0, 8)},
		{list1: listOf(0), list2: listOf(0), result: listOf(0)},
		{list1: listOf(9, 9, 9, 9, 9, 9, 9), list2: listOf(9, 9, 9, 9), result: listOf(8, 9, 9, 9, 0, 0, 0, 1)},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			result := leetcode.AddTwoNumbers(tt.list1, tt.list2)
			assert[E](t, result.Equal(tt.result))
		})
	}
}

// 0003. Longest Substring Without Repeating Characters [Medium]
func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input  string
		length int
	}{
		{input: "pwwkew", length: 3},
		{input: "abcabcbb", length: 3},
		{input: "bbbbb", length: 1},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			length := leetcode.LengthOfLongestSubstring(tt.input)
			equal[E](t, length, tt.length)
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

// 0021. Merge Two Sorted Lists [Easy]
func TestMergeTwoLists(t *testing.T) {
	listOf := linkedlist.New[int]

	tests := []struct {
		list1  *linkedlist.Node[int]
		list2  *linkedlist.Node[int]
		result *linkedlist.Node[int]
	}{
		{list1: listOf(1, 2, 4), list2: listOf(1, 3, 4), result: listOf(1, 1, 2, 3, 4, 4)},
		{list1: nil, list2: nil, result: nil},
		{list1: nil, list2: listOf(0), result: listOf(0)},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			result := leetcode.MergeTwoLists(tt.list1, tt.list2)
			assert[E](t, result.Equal(tt.result))
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

// 0028. Implement strStr() [Easy]
func TestIndexOf(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		index    int
	}{
		{haystack: "hello", needle: "ll", index: 2},
		{haystack: "aaaaa", needle: "bba", index: -1},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			index := leetcode.IndexOf(tt.haystack, tt.needle)
			equal[E](t, index, tt.index)
		})
	}
}

// 0035. Search Insert Position [Easy]
func TestSearchInsert(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		index   int
	}{
		{numbers: []int{1, 3, 5, 6}, target: 5, index: 2},
		{numbers: []int{1, 3, 5, 6}, target: 2, index: 1},
		{numbers: []int{1, 3, 5, 6}, target: 7, index: 4},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			index := leetcode.SearchInsert(tt.numbers, tt.target)
			equal[E](t, index, tt.index)
		})
	}
}

// 0053. Maximum Subarray [Easy]
func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		numbers []int
		sum     int
	}{
		{numbers: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, sum: 6},
		{numbers: []int{1}, sum: 1},
		{numbers: []int{5, 4, -1, 7, 8}, sum: 23},
	}

	for i, tt := range tests {
		t.Run(name(i), func(t *testing.T) {
			sum := leetcode.MaxSubArray(tt.numbers)
			equal[E](t, sum, tt.sum)
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

func assert[T E | F](t T, expr bool) {
	(*testing.T)(t).Helper()
	if !expr {
		f(t)("assertion failed")
	}
}

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
