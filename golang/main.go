package main

import (
	"github.com/riphidon/algoplay/golang/generics"
)

func main() {
	queue := generics.Queue[string]()
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")
	queue.Push("d")
	queue.Print()
	queue.Unshift()
	queue.Unshift()
	queue.Print()

	stack := generics.Stack[int]()
	stack.Push(12)
	stack.Push(3)
	stack.Push(9)
	stack.Push(5)
	stack.Push(8)
	stack.Push(10)
	stack.Print()
}
