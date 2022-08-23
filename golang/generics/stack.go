package generics

import "fmt"

// import "sync"

type stack[T string | int] struct {
	Size  func() int
	Push  func(T)
	Pop   func() T
	Print func()
}

func Stack[T string | int]() *stack[T] {
	// var lock *sync.Mutex
	items := []T{}
	return &stack[T]{
		Size: func() int {
			return len(items)
		},
		Push: func(data T) {
			// lock.Lock()
			// defer lock.Unlock()
			items = append(items, data)
		},
		Pop: func() T {
			// lock.Lock()
			// defer lock.Unlock()
			elt := items[len(items)-1]
			items = items[:len(items)-1]
			return elt
		},
		Print: func() {
			for _, v := range items {
				fmt.Println(v)
			}
		},
	}

}
