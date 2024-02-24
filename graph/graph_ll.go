package graph

import (
	"fmt"
)

// Node represents a node in the adjacency list
type Node struct {
	val int
	weight int
	next *Node
}

// undirected Graph_ll type using adj list
type Graph_ll struct {
	numNodes int
	adjList []*Node
}

type Digraph_ll = Graph_ll

// Define a method for the original type
func (g *Graph_ll) PrintGraph() {
	fmt.Println("Number of nodes:", g.numNodes)
	fmt.Println("Adjacency list:")
	for i, node := range g.adjList {
		fmt.Printf("%d: ",i)
		for node != nil {
			fmt.Printf("%d -> ", node.val)
			node = node.next
		}
		fmt.Printf("nil\n")
	}
}


// Create a new Graph_ll with n nodes
func CreateGraph_ll(V int) *Graph_ll {
	adjlist := make([]*Node, V)
	return &Graph_ll {
		numNodes: V,
		adjList: adjlist,
	}
}

// Create a new graph with n nodes
func CreateDigraph_ll(V int) *Digraph_ll {
	return CreateGraph_ll(V)
}

func CreateGraph_ll_m(adjm [][]int, V int) *Graph_ll {
	g := CreateGraph_ll(V)
	for i:=0;i<len(adjm);i++ {
		for j:=0;j<len(adjm[i]);j++ {
			if adjm[i][j] != 0 {
				g.AddEdge_digraph_weight(i, j, adjm[i][j])
			}
		}
	}
	return g
}

// Add an edge from node u to node v, and v to u as it is undirected
func (g *Graph_ll) AddEdge(u, v int) {
	adjlist_u := g.adjList[u]
	// new node
	node := &Node{
		val: v,
		weight: 1,
		next: nil,
	}
	node.next = adjlist_u // add on top of the list to save time
	g.adjList[u] = node // reassign this node head to adjlist
	
	adjlist_v := g.adjList[v]
	// new node again
	node = &Node{
		val: u,
		weight: 1,
		next: nil,
	}
	node.next = adjlist_v
	g.adjList[v] = node // reassign this node head to adjlist
}

// Add an edge from node u to node v, and v to u as it is undirected
func (g *Graph_ll) AddEdge_weight(u, v, w int) {
	adjlist_u := g.adjList[u]
	node := &Node{
		val: v,
		weight: w,
		next: nil,
	}
	node.next = adjlist_u // add on top of the list to save time
	g.adjList[u] = node // reassign this node head to adjlist
	
	adjlist_v := g.adjList[v]
	// new node again
	node = &Node{
		val: u,
		weight: 1,
		next: nil,
	}
	node.next = adjlist_v
	g.adjList[v] = node // reassign this node head to adjlist
}

// Add an edge from node u to node v
func (g *Graph_ll) AddEdge_digraph(u, v int) {
	adjlist_u := g.adjList[u]
	node := &Node{
		val: v,
		weight: 1,
		next: nil,
	}
	node.next = adjlist_u // add on top of the list
	g.adjList[u] = node // reassign this node head to adjlist
}

// Add an edge from node u to node v, and v to u as it is undirected
func (g *Graph_ll) AddEdge_digraph_weight(u, v, w int) {
	adjlist_u := g.adjList[u]
	node := &Node{
		val: v,
		weight: w,
		next: nil,
	}
	node.next = adjlist_u // add on top of the list
	g.adjList[u] = node // reassign this node head to adjlist
}

func (g *Graph_ll) Get_adjacent(u int) (result []int) {
	nodehead := g.adjList[u]
	for nodehead != nil {
		result = append(result, nodehead.val)
		nodehead = nodehead.next
	}
	return
}

func (g *Graph_ll) Get_edge_w(u, v int) int {
	nodehead := g.adjList[u]
	for nodehead != nil {
		if nodehead.val==v {
			return nodehead.weight
		}
		nodehead = nodehead.next
	}
	return 0
}

func (g *Graph_ll) Get_numNodes() int {
	return g.numNodes
}

func (g *Graph_ll) Get_EdgeList() (edges []Edge) {
	for u:=0;u<g.numNodes;u++ {
		nodehead := g.adjList[u]
		for nodehead != nil {
			edges = append(edges, Edge{
					u: u,
					v: nodehead.val,
					weight: nodehead.weight,
				})
			nodehead = nodehead.next
		}
	}
	return
}
