package linkedlist_test

import (
	"reflect"
	"testing"

	"github.com/junk1tm/leetcode/linkedlist"
)

func TestNew(t *testing.T) {
	var values []int
	for l := linkedlist.New(1, 2, 3); l != nil; l = l.Next {
		values = append(values, l.Value)
	}
	equal[E](t, values, []int{1, 2, 3})
}

func TestNode_String(t *testing.T) {
	list := linkedlist.New(1, 2, 3)
	equal[E](t, list.String(), "[1 2 3]")
}

func TestNode_Equal(t *testing.T) {
	tests := []struct {
		name   string
		list1  *linkedlist.Node[int]
		list2  *linkedlist.Node[int]
		result bool
	}{
		{
			name:   "equal",
			list1:  linkedlist.New(1, 2, 3),
			list2:  linkedlist.New(1, 2, 3),
			result: true,
		},
		{
			name:   "not equal",
			list1:  linkedlist.New(1, 2, 3),
			list2:  linkedlist.New(1, 2, 4),
			result: false,
		},
		{
			name:   "different length",
			list1:  linkedlist.New(1, 2, 3),
			list2:  linkedlist.New(1, 2, 3, 4),
			result: false,
		},
		{
			name:   "both nil",
			list1:  nil,
			list2:  nil,
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.list1.Equal(tt.list2)
			equal[E](t, result, tt.result)
		})
	}
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
