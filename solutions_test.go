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
