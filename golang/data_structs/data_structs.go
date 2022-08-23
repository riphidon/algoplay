package data_structs

type Item struct {
	value interface{}
}

type Stack struct {
	items []*Item
	count int
}

func NewStack() *Stack {
	return &Stack{
		items: []*Item{},
		count: 0,
	}
}

func (s *Stack) Pop() *Item {
	e := (s.items)[s.count]
	s.items = s.items[0:s.count]
	s.count--
	return e
}

func (s *Stack) Push(data *Item) {
	s.items = append(s.items, data)
	s.count++
}

type QueueNode struct {
	value interface{}
	next  *QueueNode
}

type Queue struct {
	front  *QueueNode
	back   *QueueNode
	length int
}

func NewQueue() *Queue {
	return &Queue{
		front:  nil,
		back:   nil,
		length: 0,
	}
}
func (q *Queue) Push(node *QueueNode) {
	if q.length == 0 {
		q.front = node
		q.back = node
	} else {
		q.back.next = node
		q.back = node
	}
	q.length++
}

func (q *Queue) Unshift() *QueueNode {
	removedNode := q.front
	if q.length == 0 {
		return nil
	}
	if q.length == 1 {
		q.back = nil
	}
	q.front = q.front.next
	q.length--
	return removedNode

}
