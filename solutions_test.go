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
