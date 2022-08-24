package main

import (
	"fmt"

	"github.com/riphidon/algoplay/golang/datastructs"
)

func main() {
	treeNodePlay()
	stackPlay()
	linkedListPlay()

}

func linkedListPlay() {
	list := datastructs.LinkedList[string]{}
	a := datastructs.NewLink[string]("ichi")
	b := datastructs.NewLink[string]("ni")
	c := datastructs.NewLink[string]("san")
	d := datastructs.NewLink[string]("yon")

	list.Append(a)
	list.Append(b)
	list.Append(c)
	list.Insert(b, d)
	list.Print()
	list.Remove((a))
	fmt.Println("---------------------------")
	list.Print()

}

func stackPlay() {
	stack := datastructs.Stack[int]()
	stack.Push(1)
	stack.Push(20)
	stack.Push(30)
	stack.Push(7)
	stack.Push(3)
	//fmt.Println(stack.Pop())
	//stack.Print()
}

func treeNodePlay() {
	a := datastructs.NewNode("a")
	b := datastructs.NewNode("b")
	c := datastructs.NewNode("c")
	d := datastructs.NewNode("d")
	e := datastructs.NewNode("e")
	f := datastructs.NewNode("f")

	a.Left = &b
	a.Right = &c
	b.Right = &e
	b.Left = &d
	c.Right = &f

	// a.BreadthFirstPrint()
	// a.DepthFirstPrint()
	// fmt.Println(datastructs.RecursiveDFS(a, "c"))
	// fmt.Println(datastructs.RecursiveDFS(a, "z"))
	// a.PreOrderPrint()
}
