package datastructs

import "fmt"

type ListNode[T any] struct {
	Data T
	Next *ListNode[T]
}

type LinkedList[T any] struct {
	Head   *ListNode[T]
	Length int
}

func NewLink[T any](data T) *ListNode[T] {
	return &ListNode[T]{
		Data: data,
		Next: nil,
	}
}

func (l *LinkedList[T]) Prepend(new *ListNode[T]) {
	next := l.Head
	l.Head = new
	l.Head.Next = next
	l.Length++
}

func (l *LinkedList[T]) Append(new *ListNode[T]) {
	if l.Head == nil {
		l.Head = new
		l.Length++
		return
	}
	cn := l.Head
	for cn.Next != nil {
		cn = cn.Next
	}

	cn.Next = new
}

func (l *LinkedList[T]) Insert(previous, new *ListNode[T]) {
	cn := l.Head
	for cn != nil {
		cnd := fmt.Sprintf("%v", cn.Data)
		prd := fmt.Sprintf("%v", previous.Data)
		if cnd == prd {
			new.Next = cn.Next
			cn.Next = new
			l.Length++
			return
		}
		cn = cn.Next
	}
}

func (l *LinkedList[T]) Remove(toRemove *ListNode[T]) {
	if ok := l.isNotEmpty(); ok {
		if l.Length == 0 {
			fmt.Println("The list is Empty!")
		}
		trd := fmt.Sprintf("%v", toRemove.Data)
		cn := l.Head

		if fmt.Sprintf("%v", cn.Data) == trd {
			l.Head = cn.Next
			l.Length--
			return
		}

		for cn != nil {
			cnd := fmt.Sprintf("%v", cn.Next.Data)
			if cnd == trd {
				cn.Next = cn.Next.Next
				l.Length--
				return
			}
			cn = cn.Next
		}
	}
}

func (l *LinkedList[T]) Reverse() {
	var previous *ListNode[T] = nil
	cn := l.Head
	for cn != nil {
		next := cn.Next
		cn.Next = previous
		previous = cn
		cn = next
	}
	l.Head = previous
}

func (l *LinkedList[T]) Print() {
	if ok := l.isNotEmpty(); ok {
		node := l.Head
		for node != nil {
			fmt.Println(node.Data)
			node = node.Next
		}
	}

}

func (l *LinkedList[T]) isNotEmpty() bool {
	if l.Length == 0 {
		fmt.Println("The list is Empty!")
		return false
	}

	return true
}
