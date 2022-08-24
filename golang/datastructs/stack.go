package datastructs

import (
	"fmt"
	"log"
)

type stack[T any] struct {
	Size  func() int
	Push  func(T)
	Pop   func() T
	Print func()
}

// Stack returns a stack type data container.
func Stack[T any]() *stack[T] {
	items := []T{}
	return &stack[T]{
		Size: func() int {
			return len(items)
		},
		Push: func(data T) {
			items = append(items, data)
		},
		Pop: func() T {
			if len(items) == 0 {
				log.Fatal("Empty stack!")
			}

			top := items[len(items)-1]
			items = items[:len(items)-1]
			return top
		},
		Print: func() {
			s := "Current stack:"
			for _, v := range items {
				s += fmt.Sprintf(" %v ", v)
			}
			fmt.Println(s)
		},
	}

}
