package generics

import "fmt"

type ListNode[T string | int] struct {
	Value T
	Next  *ListNode[T]
}

type queue[T string | int] struct {
	Front   func() *ListNode[T]
	Back    func() *ListNode[T]
	Push    func(T)
	Unshift func() *T
	Print   func()
}

func Queue[T string | int]() *queue[T] {
	var front *ListNode[T] = nil
	var back *ListNode[T] = nil
	length := 0
	return &queue[T]{
		Front: func() *ListNode[T] {
			return front
		},
		Back: func() *ListNode[T] {
			return back
		},
		Push: func(data T) {
			node := &ListNode[T]{
				Value: data,
				Next:  nil,
			}
			if length == 0 {
				front = node
				back = node
			} else {
				back.Next = node
				back = node
			}
			length++
		},
		Unshift: func() *T {
			removed := front
			if length == 0 {
				return nil
			}
			if length == 1 {
				back = nil
			}
			front = front.Next
			length--
			return &removed.Value
		},
		Print: func() {
			if length > 0 {
				q := fmt.Sprintf("%v", front.Value)
				node := front.Next
				for node != nil {
					q += "-->" + fmt.Sprintf("%v", node.Value)
					node = node.Next
				}
				fmt.Println(q)
			} else {
				fmt.Println("Queue is empty")
			}

		},
	}
}
