package generics

import (
	"fmt"
	"strconv"
)

type value interface {
	string | int
}

type arr[T value] []node[T]

func (a *arr[T]) shift() node[T] {
	elt := (*a)[0]
	*a = (*a)[1:]
	return elt
}

func (a *arr[T]) pop() node[T] {
	i := len((*a)) - 1
	elt := (*a)[i]
	*a = (*a)[0:i]
	return elt
}

/***********************************
*          	  NODE
***********************************/

type node[T value] struct {
	value     T
	edgesList []*node[T]
}

func newNode[T value](val T) node[T] {
	return node[T]{value: val, edgesList: []*node[T]{}}
}

func connect[T value](node *node[T], new *node[T]) {
	node.edgesList = append(node.edgesList, new)
	new.edgesList = append(new.edgesList, node)
}

func (n *node[T]) findAdjacents() {
	adj := []T{}
	for _, v := range n.edgesList {
		adj = append(adj, v.value)
	}
	fmt.Println(adj)
}

func (n *node[T]) printEdges() []T {
	edges := []T{}
	for _, v := range n.edgesList {
		edges = append(edges, v.value)
	}
	return edges
}

/***********************************
*          	 GRAPH
***********************************/
type graph[T value] struct {
	nodes []node[T]
}

func (g *graph[T]) print() {
	for _, v := range g.nodes {
		edges := v.printEdges()
		fmt.Println(v.value, edges)
	}
}

func (g *graph[T]) add(node node[T]) {
	g.nodes = append(g.nodes, node)
}

func (g *graph[T]) sum() int {
	sum := 0
	for _, node := range g.nodes {
		num, err := strconv.Atoi(fmt.Sprintf("%v", node.value))
		if err != nil {
			fmt.Println("Wrong type")
		}
		sum += num
	}
	return sum
}

func (g *graph[T]) breadthFirstPrint() {
	q := arr[T]{g.nodes[0]}
	visited := map[T]bool{
		g.nodes[0].value: true,
	}

	for len(q) > 0 {
		n := q.shift()
		for _, e := range n.edgesList {
			if !visited[e.value] {
				q = append(q, *e)

				visited[e.value] = true
			}

		}
		fmt.Println(n.value)
	}

}

func recursiveBreadthFirst[T value](node node[T], visited map[T]bool) {
	if visited == nil {
		visited = map[T]bool{}
	}

	visited[node.value] = true

	fmt.Println("visiting-->", node.value)
	for _, edge := range node.edgesList {
		if !visited[edge.value] {

			visited[edge.value] = true
			recursiveBreadthFirst(*edge, visited)
		}
	}
}

func (g *graph[T]) depthFirstPrint() {
	stack := arr[T]{g.nodes[0]}
	visited := map[T]bool{
		g.nodes[0].value: true,
	}

	for len(stack) > 0 {
		node := stack.pop()
		for _, edge := range node.edgesList {
			if !visited[edge.value] {
				stack = append(stack, *edge)
				visited[edge.value] = true
			}
		}
		fmt.Println(node.value)
	}

}

func shortestPath[T value](start, end node[T]) {
	queue := arr[T]{start}
	visited := map[T]*T{
		start.value: nil,
	}

	for len(queue) > 0 {
		node := queue.shift()
		if node.value == end.value {
			reconstructPath(visited, end)
			return
		}

		for _, edge := range node.edgesList {
			if _, ok := visited[edge.value]; !ok {
				visited[edge.value] = &node.value
				queue = append(queue, *edge)
			}
		}
	}

}

func reconstructPath[T value](nodeList map[T]*T, endNode node[T]) {
	current := &endNode.value
	pathStr := ""
	for current != nil {
		previous := current
		if pathStr == "" {
			pathStr = fmt.Sprintf("%v", *previous)
		} else {
			pathStr = fmt.Sprintf("%v", *previous) + "-->" + pathStr
		}
		current = nodeList[*previous]

	}
	fmt.Println(pathStr)
}

func main() {
	gi := graph[int]{
		nodes: []node[int]{},
	}

	nodeA := newNode(8)
	nodeB := newNode(34)
	nodeC := newNode(25)
	nodeD := newNode(5)
	nodeE := newNode(12)

	connect(&nodeA, &nodeB)
	connect(&nodeA, &nodeD)
	connect(&nodeB, &nodeC)
	connect(&nodeC, &nodeD)
	connect(&nodeD, &nodeE)

	gi.add(nodeA)
	gi.add(nodeB)
	gi.add(nodeC)
	gi.add(nodeD)
	gi.add(nodeE)

	// fmt.Println(gi.sum())
	// nodeA.findAdjacents()

	DFW := newNode("DFW")
	JFK := newNode("JFK")
	LAX := newNode("LAX")
	HNL := newNode("HNL")
	SAN := newNode("SAN")
	EWR := newNode("EWR")
	BOS := newNode("BOS")
	MIA := newNode("MIA")
	MCO := newNode("MCO")
	PBI := newNode("PBI")

	connect(&DFW, &LAX)
	connect(&DFW, &JFK)
	connect(&LAX, &HNL)
	connect(&LAX, &EWR)
	connect(&LAX, &SAN)
	connect(&JFK, &BOS)
	connect(&JFK, &MIA)
	connect(&MIA, &MCO)
	connect(&MIA, &PBI)
	connect(&MCO, &PBI)

	gs := graph[string]{
		nodes: []node[string]{
			DFW,
			JFK,
			LAX,
			HNL,
			SAN,
			EWR,
			BOS,
			MIA,
			MCO,
			PBI,
		},
	}

	gs.print()
	// gs.depthFirstPrint()
	// recursiveBreadthFirst(DFW, nil)
	shortestPath(DFW, MIA)
}
