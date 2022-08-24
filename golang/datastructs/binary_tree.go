package datastructs

import (
	"fmt"
)

type TreeNode[T any] struct {
	Value T
	Right *TreeNode[T]
	Left  *TreeNode[T]
}

// NewNode returns a TreeNode struct
// of type T. Right & Left properties set to nil.
func NewNode[T any](v T) TreeNode[T] {
	return TreeNode[T]{
		Value: v,
		Right: nil,
		Left:  nil,
	}
}

// BreadthFirstPrint explores all nodes at present depth
// and prints them before moving to nodes at next depth level.
func (n TreeNode[T]) BreadthFirstPrint() {
	queue := Queue[TreeNode[T]]()
	queue.Enqueue(n)

	for queue.Len() > 0 {
		node := queue.Dequeue()
		fmt.Println(node.Value)
		if node.Left != nil {
			queue.Enqueue(*node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(*node.Right)
		}
	}
}

// BreadthFirstSearch search the binary tree for a given value
// by searching nodes depth level by depth level.
// returns true if target is found, else returns false.
func BreadthFirstSearch[T comparable](root TreeNode[T], target T) bool {
	if root.Value == target {
		return true
	}
	queue := Queue[TreeNode[T]]()
	queue.Enqueue(root)

	for queue.Len() > 0 {
		node := queue.Dequeue()
		if node.Value == target {
			return true
		}
		if node.Left != nil {
			queue.Enqueue(*node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(*node.Right)
		}
	}

	return false
}

// DepthFirstPrint Explore nodes along each branch and prints them
// before backtraking to the next branck.
// returns true if target is found, else returns false.
func (n TreeNode[T]) DepthFirstPrint() {
	stack := Stack[TreeNode[T]]()
	stack.Push(n)

	for stack.Size() > 0 {
		node := stack.Pop()
		fmt.Println(node.Value)
		if node.Right != nil {
			stack.Push(*node.Right)
		}
		if node.Left != nil {
			stack.Push(*node.Left)
		}
	}
}

// DepthFirstSearch search the binary tree for a given value
// by searching nodes along an entire branch befor backtracking to
// the next branch.
func DepthFirstSearch[T comparable](root TreeNode[T], target T) bool {
	if root.Value == target {
		return true
	}
	stack := Stack[TreeNode[T]]()
	stack.Push(root)

	for stack.Size() > 0 {
		node := stack.Pop()
		if node.Value == target {
			return true
		}
		if node.Right != nil {
			stack.Push(*node.Right)
		}
		if node.Left != nil {
			stack.Push(*node.Left)
		}
	}
	return false
}

// RecursiveDFS implements DepthFirstSearch method in a recursive
// manner.
func RecursiveDFS[T comparable](root TreeNode[T], target T) bool {
	if root.Value == target {
		return true
	} else if root.Right != nil {
		return RecursiveDFS(*root.Right, target)
	} else if root.Left != nil {
		return RecursiveDFS(*root.Left, target)
	}
	return false
}

func (n *TreeNode[T]) PreOrderPrint() {
	root := n
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	root.Left.PreOrderPrint()
	root.Right.PreOrderPrint()
}

func (n *TreeNode[T]) PostOrderPrint() {
	root := n
	if root == nil {
		return
	}
	root.Left.PostOrderPrint()
	root.Right.PostOrderPrint()
	fmt.Println(root.Value)
}

func (n *TreeNode[T]) InOrderPrint() {
	root := n
	if root == nil {
		return
	}
	root.Left.PreOrderPrint()
	fmt.Println(root.Value)
	root.Right.PreOrderPrint()
}

func SumTree(node *TreeNode[int]) int {
	if node == nil {
		return 0
	}
	return node.Value + SumTree(node.Left) + SumTree(node.Right)
}
