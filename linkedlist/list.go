package linkedlist

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func New[T comparable](values ...T) *Node[T] {
	list := new(Node[T])
	temp := list

	for _, value := range values {
		list.Next = &Node[T]{Value: value, Next: nil}
		list = list.Next
	}

	return temp.Next
}

func (n *Node[T]) String() string {
	var values []T
	for ; n != nil; n = n.Next {
		values = append(values, n.Value)
	}

	return fmt.Sprintf("%v", values)
}

func (n *Node[T]) Equal(n2 *Node[T]) bool {
	for ; n != nil && n2 != nil; n, n2 = n.Next, n2.Next {
		if n.Value != n2.Value {
			return false
		}
	}
	if n != nil || n2 != nil {
		return false
	}

	return true
}
