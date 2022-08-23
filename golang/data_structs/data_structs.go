package data_structs

type Item struct {
	Value interface{}
}

type Stack struct {
	Items []*Item
	Count int
}

func NewStack() *Stack {
	return &Stack{
		Items: []*Item{},
		Count: 0,
	}
}

func (s *Stack) Pop() *Item {
	e := (s.Items)[s.Count]
	s.Items = s.Items[0:s.Count]
	s.Count--
	return e
}

func (s *Stack) Push(data *Item) {
	s.Items = append(s.Items, data)
	s.Count++
}

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type Queue struct {
	Front  *ListNode
	Back   *ListNode
	Length int
}

func NewQueue() *Queue {
	return &Queue{
		Front:  nil,
		Back:   nil,
		Length: 0,
	}
}
func (q *Queue) Push(value interface{}) {
	node := &ListNode{
		Value: value,
		Next:  nil,
	}
	if q.Length == 0 {
		q.Front = node
		q.Back = node
	} else {
		q.Back.Next = node
		q.Back = node
	}
	q.Length++
}

func (q *Queue) Unshift() *ListNode {
	removedNode := q.Front
	if q.Length == 0 {
		return nil
	}
	if q.Length == 1 {
		q.Back = nil
	}
	q.Front = q.Front.Next
	q.Length--
	return removedNode

}
