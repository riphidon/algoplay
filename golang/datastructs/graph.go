package datastructs

import (
	"fmt"
	"strconv"
)

type value interface {
	string | int
}

type arr[T value] []GraphNode[T]

func (a *arr[T]) shift() GraphNode[T] {
	elt := (*a)[0]
	*a = (*a)[1:]
	return elt
}

func (a *arr[T]) pop() GraphNode[T] {
	i := len((*a)) - 1
	elt := (*a)[i]
	*a = (*a)[0:i]
	return elt
}

/***********************************
*          	  NODE
***********************************/

type GraphNode[T value] struct {
	value     T
	edgesList []*GraphNode[T]
}

func newNode[T value](val T) GraphNode[T] {
	return GraphNode[T]{value: val, edgesList: []*GraphNode[T]{}}
}

func (n *GraphNode[T])Connect(new *GraphNode[T]) {
	n.edgesList = append(n.edgesList, new)
	new.edgesList = append(new.edgesList, n)
}

func (n *GraphNode[T]) findAdjacents() {
	adj := []T{}
	for _, v := range n.edgesList {
		adj = append(adj, v.value)
	}
	fmt.Println(adj)
}

func (n *GraphNode[T]) printEdges() []T {
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
	nodes []GraphNode[T]
}

func (g *graph[T]) print() {
	for _, v := range g.nodes {
		edges := v.printEdges()
		fmt.Println(v.value, edges)
	}
}

func (g *graph[T]) add(node GraphNode[T]) {
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

func recursiveBreadthFirst[T value](node GraphNode[T], visited map[T]bool) {
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

func shortestPath[T value](start, end GraphNode[T]) {
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

func reconstructPath[T value](nodeList map[T]*T, endNode GraphNode[T]) {
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

func graphPlay() {
	gi := graph[int]{
		nodes: []GraphNode[int]{},
	}

	nodeA := newNode(8)
	nodeB := newNode(34)
	nodeC := newNode(25)
	nodeD := newNode(5)
	nodeE := newNode(12)

	nodeA.Connect(&nodeB)
	nodeA.Connect(&nodeD)
	nodeB.Connect(&nodeC)
	nodeC.Connect(&nodeD)
	nodeD.Connect(&nodeE)

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

	DFW.Connect(&LAX)
	DFW.Connect(&JFK)
	LAX.Connect(&HNL)
	LAX.Connect(&EWR)
	LAX.Connect(&SAN)
	JFK.Connect(&BOS)
	JFK.Connect(&MIA)
	MIA.Connect(&MCO)
	MIA.Connect(&PBI)
	MCO.Connect(&PBI)

	gs := graph[string]{
		nodes: []GraphNode[string]{
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
