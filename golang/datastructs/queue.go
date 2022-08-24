package datastructs

import "fmt"

type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}

type queue[T any] struct {
	Front   func() *QueueNode[T]
	Back    func() *QueueNode[T]
	Enqueue func(T)
	Dequeue func() *T
	Len     func() int
	Print   func()
}

// Queue returns a queue type object container.
func Queue[T any]() *queue[T] {
	var front *QueueNode[T] = nil
	var back *QueueNode[T] = nil
	length := 0
	return &queue[T]{
		Front: func() *QueueNode[T] {
			return front
		},
		Back: func() *QueueNode[T] {
			return back
		},
		Enqueue: func(data T) {
			node := &QueueNode[T]{
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
		Dequeue: func() *T {
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
		Len: func() int {
			return length
		},
	}
}
